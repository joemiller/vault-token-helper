package cmd

import "github.com/spf13/viper"

// config file example, yaml:
//
//   # $HOME/.vault-token-helper.yaml
//   ---
//   backend: [automatic, (default)
//             keychain,
//             secret-service,
//             pass,
//             wincred]
//
//   keychain:
//    keychain: (optional) name of macOS keychain to use. Default is the standard login keychain
//    icloud: (optional) boolean. If true, the secret will be syncronizable to iCloud Keychain. Default: false
//
//   secret_service:
//     collection: (optional) name of the collection to store tokens in (default: 'vault')
//
//   pass:
//     dir: (optional) path to password-store directory
//     command: (optional) path to the pass executable
//     prefix: (optional) prefix to apply to stored tokens (default: 'vault')
//
//   wincred:
//     prefix: (optional) prefix to apply to stored tokens (default: 'vault')
//
type config struct {
	BackendType   string                     `mapstructure:"backend"`
	Keychain      keychainBackendConfig      `mapstructure:"keychain"`
	SecretService secretServiceBackendConfig `mapstructure:"secret-service"`
	Pass          passBackendConfig          `mapstructure:"pass"`
	WinCred       wincredBackendConfig       `mapstructure:"wincred"`
}

type keychainBackendConfig struct {
	Keychain       string `mapstructure:"keychain_name"`
	Synchronizable bool   `mapstructure:"icloud"`
}

type secretServiceBackendConfig struct {
	Collection string `mapstructure:"collection"`
}

type passBackendConfig struct {
	Dir     string `mapstructure:"dir"`
	Command string `mapstructure:"command"`
	Prefix  string `mapstructure:"prefix"`
}

type wincredBackendConfig struct {
	Prefix string `mapstructure:"prefix"`
}

func initDefaultConfig() {
	viper.SetDefault("Keychain",
		keychainBackendConfig{
			Synchronizable: false,
		},
	)

	viper.SetDefault("SecretService",
		secretServiceBackendConfig{
			Collection: "vault",
		},
	)

	viper.SetDefault("Pass",
		passBackendConfig{
			Prefix: "vault",
		},
	)

	viper.SetDefault("WinCred",
		wincredBackendConfig{
			Prefix: "vault",
		},
	)
}
