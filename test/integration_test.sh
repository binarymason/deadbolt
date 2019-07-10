#!/usr/bin/env bash

set -e

pwd

test_secret=foo
url=localhost:8080

setup() {
  go install
  echo "# starting deadbolt..."
  DEADBOLT_SECRET="$test_secret" deadbolt -c ./testdata/simple_deadbolt_config.yml &>/dev/null &
  sleep 3
}

fail() {
  echo "FAIL! $*"
  exit 1
}

ok() {
  echo "--> OK"
}

cleanup() {
  echo "# killing server"
  pkill -f "deadbolt"
}
trap cleanup EXIT

setup


echo "# test root endpoint"
curl -f "$url"
ok


echo "# test 404 endpoint"
status_code=$(curl -sw '%{http_code}' "$url/fooobar" | tail -n 1)
if [ "$status_code" != "404" ]; then
  fail "expected status code for an unknown endpoint to be 404 but got $status_code"
fi
ok


echo "# test unauthorized /lock"
status_code=$(curl -sw '%{http_code}' -XPOST "$url/lock" | tail -n 1)
if [ "$status_code" != "401" ]; then
  fail "expected status code to be 401 but got $status_code"
fi
ok

echo "# test unauthorized /unlock"
status_code=$(curl -sw '%{http_code}' -XPOST "$url/unlock" | tail -n 1)
if [ "$status_code" != "401" ]; then
  fail "expected status code to be 401 but got $status_code"
fi
ok

echo "# test /lock"
status_code=$(curl -sw '%{http_code}' -XPOST -H "Authorization: $test_secret" "$url/lock" | tail -n 1)
if [ "$status_code" != "201" ]; then
  fail "expected status code to be 201 but got $status_code"
fi
ok

echo "# test /unlock"
status_code=$(curl -sw '%{http_code}' -XPOST -H "Authorization: $test_secret" "$url/unlock" | tail -n 1)
if [ "$status_code" != "201" ]; then
  fail "expected status code to be 201 but got $status_code"
fi
ok

# TODO: when writing to files, test that it actually worked :)

echo "PASS"
