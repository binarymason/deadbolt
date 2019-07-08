> WIP

# Simple HTTP server to toggle PermitRootLogin sshd settings remotely

Why in the world would you want to do this?  Permitting root access is typically a security opening.  But, if you _must_ remote log into a server as root, this adds a layer of security. In order for deadbolt to accept your request, you must be an authorized IP address and pass in the correct `Authorization` header.


### Development: Running tests

If you have inotify-tools installed, you can start this script and tests will be automatically run on any file change.
```
./script/watch_tests.sh
```
