module github.com/joemiller/vault-token-helper

go 1.17

require (
	github.com/99designs/keyring v0.0.0-20190704105226-2c916c935b9f
	github.com/PuerkitoBio/purell v1.1.0
	github.com/hashicorp/vault/api v1.0.2
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.7.5
)

require (
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/danieljoos/wincred v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dvsekhvalnov/jose2go v0.0.0-20180829124132-7f401d37b68a // indirect
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.5.3 // indirect
	github.com/hashicorp/go-rootcerts v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/vault/sdk v0.1.8 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/keybase/go-keychain v0.0.0-20190604185112-cc436cc9fe98 // indirect
	github.com/magiconair/properties v1.8.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/spf13/afero v1.1.2 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sys v0.0.0-20190626221950-04f50cda93cb // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	gopkg.in/square/go-jose.v2 v2.3.1 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e => ../keyring
// replace github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e => github.com/joemiller/keyring v0.0.0-20190624143912-6409680b37b7b84fe91df0532f82861e9e4343c8
// replace github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c => ../go-libsecret
