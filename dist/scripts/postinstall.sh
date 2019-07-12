#!/usr/bin/env bash

service_manager=$(ps --no-headers -o comm 1)

if [ "$service_manager" = init ]; then
  echo "+ setting up init script"
  mv /etc/deadbolt/init/sysvinit/deadbolt /etc/init.d/deadbolt
else
  echo "+ setting up systemd service"
  mv /etc/deadbolt/init/systemd/deadbolt.service /etc/systemd/system/deadbolt.service
fi

echo "+ cleanup"
rm -rf "/etc/deadbolt/init"

echo "+ DONE"
