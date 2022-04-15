Audisp-go
=========

[![Build Status](https://drone.garambrogne.net/api/badges/athoune/audisp-go/status.svg)](https://drone.garambrogne.net/athoune/audisp-go)
[![go-report](https://goreportcard.com/badge/github.com/athoune/audisp-go)](https://goreportcard.com/report/github.com/athoune/audisp-go)

[Godoc](https://pkg.go.dev/github.com/athoune/audisp-go)

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
    ./bin/  audisp-expr 'line.type == "SYSCALL" and line.syscall == syscall("connect") and line.comm == "curl" '

Do something that trigger auditd, some curl
