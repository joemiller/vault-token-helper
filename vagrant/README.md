vagrant
=======

This directory contains Vagrant images for use in development and testing.

Using
-----


### Fedora (Gnome desktop)

**Pre-Reqs:**

1. Install vagrant
2. Install vagrant plugins:

```
vagrant plugin install vagrant-vbguest
vagrant plugin install vagrant-reload
```

**Launch**:

```
cd vagrant/fedora
vagrant up
```

This will launch the fedora VM with a Gnome UI and Gnome Keyring installed.
A full go environment will also be installed. The first `up` may take a while
to install all the packages and reboot at least once. You may want to use
`vagrant halt` instead of `destroy` to suspend the VM until you're done
with dev/test.

The root of the project will be mounted as a host folder to `/vault-token-helper`.

Run tests from an SSH or GUI Terminal session in the fedora VM:

```
cd /vault-token-helper
make test
# or:
go test -v ./...
```

> If you run the tests via ssh you need to login to the web UI first in order for
> the gnome keyring service to start first.

### Windows 10

**Pre-Reqs:**

1. Install vagrant

**Launch**:

```
cd vagrant/windows
vagrant up
```

`git` and `go` will be installed via the chocolately package manager.

A GUI will open up. Login and open cmd or powershell.

The root of the project will be mounted to `C:\vault-token-helper`
