#!/usr/bin/env bash

service_manager=$(ps --no-headers -o comm 1)

banner() {
  echo "##################################"
}

common_instructions() {
  echo "REQUIRED: add authorized_keys to /etc/deadbolt/deadbolt.yml"
}

if [ "$service_manager" = init ]; then
  echo "+ setting up init script"
  mv /etc/deadbolt/init/sysvinit/deadbolt /etc/init.d/deadbolt
  banner
  echo "run the following to get started:"
  common_instructions
  echo "$ chkconfig --add deadbolt"
  echo "$ service deadbolt start"
  banner
else
  echo "+ setting up systemd service"
  mv /etc/deadbolt/init/systemd/deadbolt.service /etc/systemd/system/deadbolt.service
  banner
  echo "run the following to get started:"
  common_instructions
  echo "$ systemctl daemon-reload"
  echo "$ systemctl enable deadbolt"
  echo "$ systemctl start deadbolt"
  banner
fi

echo "+ cleanup"
rm -rf "/etc/deadbolt/init"

echo "+ DONE"
