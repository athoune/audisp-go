Audisp-go
=========

`audisp` client for Linux auditd `service`.

Test it
-------

Edit your `audisp` `af_unix` config

    vi /etc/audisp/plugins.d/af_unix.conf

```
# This file controls the configuration of the
# af_unix socket plugin. It simply takes events
# and writes them to a unix domain socket. This
# plugin can take 2 arguments, the path for the
# socket and the socket permissions in octal.

active = yes
direction = out
path = builtin_af_unix
type = builtin
args = 0640 /var/run/audispd_events
format = string
```

`active = yes` and `args` path are important.

You can now build and test:

    make
    ./bin/audisp

Do something that trigger auditd, like ssh connection, sudo something.
