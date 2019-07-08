module github.com/joemiller/vault-token-helper

require (
	github.com/99designs/keyring v0.0.0-20190704105226-2c916c935b9f
	github.com/PuerkitoBio/purell v1.1.0
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/hashicorp/vault/api v1.0.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	google.golang.org/appengine v1.4.0 // indirect
)

// replace github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e => ../keyring
// replace github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e => github.com/joemiller/keyring v0.0.0-20190624143912-6409680b37b7b84fe91df0532f82861e9e4343c8
// replace github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c => ../go-libsecret
