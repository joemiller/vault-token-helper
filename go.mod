module github.com/joemiller/vault-token-helper

require (
	github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e
	github.com/PuerkitoBio/purell v1.1.0
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.0
	github.com/sebdah/goldie v0.0.0-20190531093107-d313ffb52c77
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
)

replace github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e => ../keyring

// replace github.com/99designs/keyring v0.0.0-20190531235905-2e3b4e59b02e => github.com/joemiller/keyring v0.0.0-20190624143912-83f5f71e04d29a7c209cd71d23c27a3f29d12178

// replace github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c => ../go-libsecret
