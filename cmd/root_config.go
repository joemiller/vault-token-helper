package cmd

import "github.com/spf13/viper"

// When modifying this file be sure to update the "vault-token-helper.annotated.yaml"
// at the root of the repo.

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
