# SwitchHosts
A small command line utility for managing your hosts file.

# Installation
Make sure your `GOPATH` binary directory is in your path:

```sh
$ export PATH=$PATH:$GOPATH/bin
```

Then install SwitchHosts:
```sh
$ go get github.com/kyledayton/switch-hosts/...
```

# Usage
```sh
Usage of switch-hosts:
switch-hosts <HOSTS_CONFIG>
<HOSTS_CONFIG> must be a file present in ~/.hosts/
Use the 'default' config to restore the original hosts file
```

The first time you run `switch-hosts`, a config directory will be created in your home folder `~/.hosts/`.
A backup of your hosts file `/etc/hosts` will be created at `/etc/hosts.orig`. In addition, a copy of your original hosts file will be stored at `~/.hosts/default`. **DO NOT REMOVE THIS FILE!** When you switch hosts, the resulting hosts file will be the `default` with your selected HOSTS_CONFIG appended to it.

### Example:
#### ~/.hosts/default
```
127.0.0.1 localhost
```

#### ~/.hosts/staging
```
10.0.0.0 staging.example.com
```

After running `switch-hosts staging`, your hosts file (`/etc/hosts`) will be:
```
127.0.0.1 localhost

10.0.0.0 staging.example.com
```

# Adding configurations
To create a configuration, simply add it to your host config directory:

```
$ touch ~/.hosts/new-config
$ $EDITOR ~/.hosts/new-config

$ switch-hosts new-config
```

# Restoring Default Configuration
To restore your default hosts file, simply use the 'default' config:
```
$ switch-hosts default
```

# License
```
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
```
