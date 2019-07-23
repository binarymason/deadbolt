This jinja file is an example of an [ansible template](https://docs.ansible.com/ansible/latest/modules/template_module.html) for setting up a [bashRPC server](https://github.com/binarymason/bashRPC) that can toggle sshd settings remotely.

```yml
# bashrpc.yml.j2
---

port: {{ bashrpc_port }}
whitelisted_clients:
{% for client in bashrpc_whitelisted_clients %}
  - {{ client }}
{% endfor %}

secret: {{ bashrpc_secret }}

routes:
  - path: /deadbolt/status
    cmd: grep PermitRootLogin /etc/ssh/sshd_config

  - path: /deadbolt/unlock
    cmd: |
      echo "{{ deadbolt_control_server_ssh_key }}" > /root/.ssh/authorized_keys
      sshd_config=/etc/sshd_config
      oldmd5="$(md5sum $sshd_config)"
      sed -ri "s|^#?PermitRootLogin .*$|PermitRootLogin without-password" "$sshd_config"
      newmd5="$(md5sum $sshd_config)"
      if [ "$newmd5" != "$oldmd5" ]; then
        service sshd restart
      fi
      echo "unlock successful"

  - path: /deadbolt/lock
    cmd: |
      rm /root/.ssh/authorized_keys
      sshd_config=/etc/sshd_config
      oldmd5="$(md5sum $sshd_config)"
      sed -ri "s|^#?PermitRootLogin .*$|PermitRootLogin no" "$sshd_config"
      newmd5="$(md5sum $sshd_config)"
      if [ "$newmd5" != "$oldmd5" ]; then
        service sshd restart
      fi
      echo "lock successful"
```
