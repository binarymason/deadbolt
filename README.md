# Simple HTTP server to toggle PermitRootLogin sshd settings remotely

Why in the world would you want to do this?  Permitting root access is typically a security opening.  But, if you _must_ remote log into a server as root, this adds a layer of security. In order for deadbolt to accept your request, you must be an authorized IP address and pass in the correct `Authorization` header.


### Configuration:
Options must be specified in a deadbolt.yml file.  The default location is `/etc/deadbolt/deadbolt.yml` but you can specify a custom location on command line.

Example deadbolt.yml

```
---

port: 8675

whitelisted_clients:
- 127.0.0.1
- 127.0.0.2

deadbolt_secret: supersecret

authorized_keys:
- from="127.0.0.1" ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCZhk9h+S5bm8gCKqbEML911qlu25iLtdFNzTJph3YlKdZV0+xJhz....pIMjmXRPf8R4NyxosTNdh79L8h6kDPe2QG6Imbo1i1H/EvgVnZ74EPcXcw1kpSsYBkyFIiqCFiC/B8+O5/APgnvNxKKNcaombnGtSiDWcVQ9DQoy/3SAI2r3AVKrYk+JVBUmxdD/Ix m@tp
- ssh-rsa AABBB3NzaC1yc2EAAAADAQABAAABAQCZhk9h+S5bm8gCKqbEML911qlu25iLtdFNzTJph3YlKdZV0+xJhz....pIMjmXRPf8R4NyxosTNdh79L8h6kDPe2QG6Imbo1i1H/EvgVnZ74EPcXcw1kpSsYBkyFIiqCFiC/B8+O5/APgnvNxKKNcaombnGtSiDWcVQ9DQoy/3SAI2r3AVKrYk+JVBUmxdD/Ix m@tp

```

While you can specify the `deadbolt_secret` in `deadbolt.yml`, you can also use the environment variable `DEADBOLT_SECRET`.


### Usage:

```
deadbolt -c /path/to/deadbolt.yml
```


### Development: Running tests

If you have inotify-tools installed, you can start this script and tests will be automatically run on any file change.
```
./script/watch_tests.sh
```
