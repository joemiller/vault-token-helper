vault-token-helper
==================

A @hashicorp Vault [token helper](https://www.vaultproject.io/docs/commands/token-helper.html) with
support of native secret storage backends on macOS, Linux, and Windows.

Features
--------

Store and retrieve tokens for multiple Vault (`$VAULT_ADDR`) instances, simplifying operators'
workflows when working with multiple Vaults.

Supported backends:

* macOS Keychain
* Linux (DBus Secret Service compatible backends, eg: Gnome Keyring)
* Windows (WinCred)
* [pass](https://www.passwordstore.org/)

Install
-------

### One-line install

| OS    | Command                                          |
| ----- | ------------------------------------------------ |
| macOS | `brew install joemiller/taps/vault-token-helper` |

### Linux package install

| Format                                                                 | Architectures |
| ---------------------------------------------------------------------- | ------------- |
| [rpm](https://github.com/joemiller/vault-token-helper/releases/latest) | amd64         |
| [deb](https://github.com/joemiller/vault-token-helper/releases/latest) | amd64         |

### Pre-built binaries

| OS      | Architectures | release                                                                               |
| ------- | ------------- | ------------------------------------------------------------------------------------- |
| macOS   | amd64         | [vault-token-helper](https://github.com/joemiller/vault-token-helper/releases/latest) |
| Linux   | amd64         | [vault-token-helper](https://github.com/joemiller/vault-token-helper/releases/latest) |
| Windows | amd64         | [vault-token-helper](https://github.com/joemiller/vault-token-helper/releases/latest) |

### From source

Clone this repo and compile for the current architecture:

```sh
make build
```

Binaries for all supported platforms are built using the [dockercore/golang-cross](https://github.com/docker/golang-cross)
image. This is the same image used by the docker cli project. The image makes it possible to
cross-compile and link to platform-specific libraries such as the OSX SDK on macOS:

```sh
make snapshot
```

Usage
-----

### Configure Vault

Install `vault-token-helper` then run:

```sh
vault-token-helper enable
```

This creates (overwrites) the `$HOME/.vault` config file with the following contents. The
`vault` CLI uses this config file to find and execute the token helper.

```toml
token_helper = "/install/path/to/vault-token-helper"
```

### Configure vault-token-helper

For most installations the defaults should be sufficient. An optional configuration file
located at `$HOME/.vault-token-helper.yaml` can be used to override the defaults.

A fully annotated example config file is available in [./vault-token-helper.annotated.yaml](./vault-token-helper.annotated.yaml)

### Login to Vault

Set `VAULT_ADDR` to the URL of your Vault instance and run `vault` commands like normal. For example,
to login and store a token on a Vault instance with the Okta auth plugin enabled:

```sh
export VAULT_ADDR=https://vault:8200
vault login -method=okta username=joe@dom.tld
```

Upon successful authentication the Vault token will be stored securely in the platform's
secrets store.

Support for storing tokens from multiple Vault instances is implemented. Change the `VAULT_ADDR`
environment variable to switch between Vault instances.

### Additional commands

The standard `store`, `get`, and `erase` commands are implemented according to the vault token
helper [spec](https://www.vaultproject.io/docs/commands/token-helper.html).

There are a few additional commands:

* `enable`: Enable the vault-token-helper by (over)writing the ~/.vault config file.
* `backends`: List the available backends on the current platform
* `list`: List tokens. Add `--extended` flag to lookup additional details about the stored
    token by quering the Vault instance's token lookup API.

Support
-------

Please open a GitHub [issue](https://github.com/joemiller/vault-token-helper/issues).

Development
-----------

### Tests

Run tests: `make test`.

Some tests are platform specific and difficult to test outside of a full desktop environment
due to interactive elements such as password prompts. To aid in development there are Vagrant
VMs with GUIs enabled in the `./vagrant/` directory. See the
[./vagrant/README.md](./vagrant/README.md) for further details.

The most complete way to run all tests would be to run `make test` running under each platform.

There is test coverage in `pkg/store` covering all of the supported backends. Additionally, there
is an integration test in the `cmd` package.

### CI/CD

Azure DevOps Pipelines is used for CI and CD because it provides support for macos, windows, and linux.
Tests are run on pull requests and releases are generated on successful master branch builds.

### Release Management

Releases are cut automatically on all successful master branch builds. This project uses
[autotag](https://github.com/pantheon-systems/autotag) and [goreleaser](https://goreleaser.com/) to
automate this process.

Semver (vMajor.Minor.Patch) is used for versioning and releases. By default, autotag will bump the
patch version on a successful master build, eg: `v1.0.0` -> `v1.0.1`.

To bump the major or minor release instead, include `[major]` or `[minor]` in the commit message.
Refer to the autotag [docs](https://github.com/pantheon-systems/autotag#incrementing-major-and-minor-versions)
for more details.

Include `[skip ci]` in the commit message to prevent a new version from being released. Only use this
for things like documentation updates.

TODO
----

*after v0.1.0:*

* [ ] The wincred lib used by 99designs/keyring has more configuration options available. Make these available in 99designs/keyring and vault-token-helper.
* [ ] add a flag like `--lookup` to `list` that will query vault for additional token info, eg: valid/invalid, ttl, policies
* ci/cd:
  * [ ] `sign` checksum.txt and assets in goreleaser.yaml GPG key
  * [ ] apple `codesign` the macos binaries
  * [ ] figure out how to cache go modules in azure pipelines, using this task maybe - https://github.com/microsoft/azure-pipelines-artifact-caching-tasks
  * [ ] linux tests, figure out how to test dbus secret-service in headless CI. probably need a stub to connect to Dbus and provide the 'prompt' service
