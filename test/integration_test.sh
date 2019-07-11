#!/usr/bin/env bash

set -e

test_secret=foo
test_sshd_config=/tmp/test_sshd_config
url=localhost:8080

cmd="deadbolt -c ./testdata/simple_deadbolt_config.yml"

setup() {
  go install

  cp ./testdata/commented_locked_sshd_config  "$test_sshd_config"
  echo "# starting deadbolt..."


  DEADBOLT_SSHD_CONFIG="$test_sshd_config" \
    DEADBOLT_SECRET="$test_secret" \
    eval "$cmd" &>/dev/null &
  sleep 3
}

fail() {
  echo "!!! $*"
  echo "FAIL"
  exit 1
}

ok() {
  echo "--> OK"
}

cleanup() {
  rm "$test_sshd_config"
  pkill -f "$cmd"
}
trap cleanup EXIT

setup


echo "# test root endpoint"
curl -f "$url"
ok


assert_status_code() {
  local expected_code="$1"
  local status_code

  status_code=$(curl -sw '%{http_code}' "${@:2}" | tail -n 1)
  if [ "$status_code" != "$expected_code" ]; then
    fail "expected status code to be $expected_code but got $status_code"
  fi

}

echo "# test 404 endpoint"
assert_status_code 404 "$url/foobar"
ok


echo "# test unauthorized /lock"
assert_status_code 401 -XPOST "$url/lock"
ok

echo "# test unauthorized /unlock"
assert_status_code 401 -XPOST "$url/unlock"
ok

assert_setting() {
  local setting
  local actual

  setting="$1"
  actual=$(grep -E '^#?PermitRootLogin' "$test_sshd_config")
  if ! echo "$actual" | grep -E "^PermitRootLogin $setting" >/dev/null; then
    fail "expected $setting but got: $actual"
  fi
}

configs=(
commented_locked_sshd_config
commented_unlocked_sshd_config
locked_sshd_config
simple_locked_sshd_config
simple_unlocked_sshd_config
unlocked_sshd_config
)

with_config() {
  local config="$1"
  echo -e "\t+ with $config"
  cp "./testdata/$config" "$test_sshd_config"
}

echo "# test /lock"
for c in "${configs[@]}"; do
  with_config "$c"
  assert_status_code 201 -XPOST -H "Authorization: $test_secret" "$url/lock"
  assert_setting no
done
ok

echo "# test /unlock"
for c in "${configs[@]}"; do
  with_config "$c"
  assert_status_code 201 -XPOST -H "Authorization: $test_secret" "$url/unlock"
  assert_setting without-password
done
ok

echo "PASS"
