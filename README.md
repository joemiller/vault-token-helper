vault-gpg-token-helper
======================

[![CircleCI](https://circleci.com/gh/joemiller/vault-gpg-token-helper.svg?style=svg)](https://circleci.com/gh/joemiller/vault-gpg-token-helper)

A @hashicorp Vault [token helper](https://www.vaultproject.io/docs/commands/token-helper.html) for storing tokens in a GPG encrypted file.

Features:

* Supports storing multiple tokens based on the `$VAULT_ADDR` env var.
* Supports GPG keys stored on YubiKey and other smartcards.

Install
-------

### Requirements

* `vault` cli (macOS: `brew install vault`)
* `gpg` (Tested with 2.2.x, likely compatible with 1.x and 2.1, macOS: `brew install gnupg`)

A `gpg` binary should be in your `$PATH`. An explicit path can be set with the
`VAULT_GPG_BIN` environment variable.

This program uses the gpg binary instead of Go's opengpg library to make it possible
to utilize GPG keys stored on a hardware device such as a YubiKey.

1. Install binary:

  * Binary releases are [available](https://github.com/joemiller/vault-gpg-token-helper/releases) for many platforms.
  * Homebrew (macOS): `brew install joemiller/taps/vault-gpg-token-helper`

2. After installation:

  * Create a `~/.vault` file with contents:

    ```toml
    token_helper = "/path/to/vault-gpg-token-helper"
    ```

    > For homebrew installations you can create this file by running:
    >
    > ```console
    > echo "token_helper = \"$(brew --prefix joemiller/taps/vault-gpg-token-helper)/bin/vault-gpg-token-helper\"" > ~/.vault
    > ```

3. Create config file `~/.vault-gpg-token-helper.toml`:

The default config file is `~/.vault-gpg-token-helper.toml`. This can be changed with the
`VAULT_GPG_CONFIG` environment variable.

At minimum a `gpg_key_id` must be set in the config file. Alternatively it can be
specified by the `VAULT_GPG_KEY_ID` environment variable.

Example:

```toml
gpg_key_id = "first last (yubikey) <firstlast@dom.tld>"
```

> Run `gpg --list-keys` for a list of keys.

Usage
-----

The `VAULT_ADDR` environment variable must be set. The storer uses this variable
as an index for storing and retrieving tokens. This allows for easy switching
between multiple Vault targets.

Example, adding a token to the store:

```console
export VAULT_ADDR="https://vault-a:8200"
vault login
```

Listing contents of the token store can be done with `gpg`, assuming you are using
the default storage path:

```console
gpg -d ~/.vault_tokens.gpg
```

> CAREFUL! Tokens will be printed in the clear to your console. In the future we may
> implement a safer 'list' command.

Creating a GPG keypair
----------------------

If you don't have a GPG key yet you can create one with:

```console
gpg --full-generate-key
```

Or, if using hardware key like a YubiKey with the OpenPGP applet enabled:

```console
gpg --card-edit

gpg/card> admin
gpg/card> generate
…
```

Token Storage
-------------

Tokens are stored encrypted in `~/.vault_tokens.gpg` by default. This can be
changed by:

* Setting the `token_db_file` configuration file option
* Setting the `VAULT_GPG_TOKEN_STORE` environment variable

Environment variables take precedence over configuration file settings.

> Vault 0.10.2+ supports a `-no-print` flag to store the token without printing to stdout

Support
-------

Please open a GitHub issue.

Release Management
------------------

Releases are cut automatically on a successful master branch build. This project uses
[autotag](https://github.com/pantheon-systems/autotag) and [goreleaser](https://goreleaser.com/) to automate this process.

Semver (vMajor.Minor.Patch) is used for versioning and releases. By default, autotag will bump the patch version
on a successful master build, eg: `v1.0.0` -> `v1.0.1`.

To bump the major or minor release instead, include the text `[major]` or `[minor]` in the commit message.
See the autotag [docs](https://github.com/pantheon-systems/autotag#incrementing-major-and-minor-versions) for more details.

To prevent a new release being built, include `[ci skip]` in the commit message. Only use this for things like documentation updtes.

TODO
----

TODOs have moved to github [issues](https://github.com/joemiller/vault-gpg-token-helper/issues)
