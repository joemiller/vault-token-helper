package cmd

import (
	"fmt"
	"os"

	"github.com/99designs/keyring"
	"github.com/joemiller/vault-token-helper/pkg/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	app = "vault-token-helper"

	cfg     config // root_config.go
	debug   bool
	cfgFile string
	backend *store.Store
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   app,
	Short: "Vault Token Helper",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	// Send all usage and error messages to STDERR. Vault expects token helpers will only
	// ever return tokens on STDOUT, all other informational messages must go through STDERR
	//	RootCmd.SetOutput(os.Stderr)
	RootCmd.SetOut(os.Stdout)
	RootCmd.SetErr(os.Stderr)

	cobra.OnInitialize(initDefaultConfig)
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vault-token-helper.yaml)")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable additional logging")
}

// initConfig reads in config file
func initConfig() {
	initDefaultConfig() // config.go

	// handle debug flag
	if debug {
		keyring.Debug = true
	}

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	// TODO: add the XDG config path standards
	viper.SetConfigName(".vault-token-helper") // name of config file (without extension)
	viper.AddConfigPath("$HOME")               // adding home directory as first search path

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Don't fail if a config file is not present. This makes it easier to implement commands like 'version' which
		// don't require a config file.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintf(os.Stderr, "Unable to read config file: %s\n", err)
			os.Exit(1)
		}
	} else {
		if err := viper.Unmarshal(&cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse config file: %s\n", err)
			os.Exit(1)
		}
	}
}

// initBackend initializes the token storage into the package var 'backend'.
// It is not called during the early init phase to avoid errors with commands
// that do not need access to a backend. Instead, commands that interact with a backend
// should call initBackend and propagate errors back to the rootcmd.
func initBackend() error {
	var err error

	storeCfg := keyring.Config{
		ServiceName: "vault",

		// keychain (macos)
		KeychainTrustApplication: true,
		KeychainSynchronizable:   cfg.Keychain.Synchronizable,

		// secret-service
		LibSecretCollectionName: cfg.SecretService.Collection,

		// pass
		PassDir:    cfg.Pass.Dir,
		PassCmd:    cfg.Pass.Command,
		PassPrefix: cfg.Pass.Prefix,

		// wincred
		WinCredPrefix: cfg.WinCred.Prefix,
	}

	switch cfg.BackendType {
	case "automatic", "":
		storeCfg.AllowedBackends = store.SupportedBackends
	case "keychain":
		storeCfg.AllowedBackends = []keyring.BackendType{keyring.KeychainBackend}
	case "secret-service":
		storeCfg.AllowedBackends = []keyring.BackendType{keyring.SecretServiceBackend}
	case "wincred":
		storeCfg.AllowedBackends = []keyring.BackendType{keyring.WinCredBackend}
	case "pass":
		storeCfg.AllowedBackends = []keyring.BackendType{keyring.PassBackend}
	default:
		return errors.Errorf("Unknown backend '%s'", cfg.BackendType)
	}

	backend, err = store.New(storeCfg)
	if err != nil {
		return errors.Wrapf(err, "Unable to initialize backend '%s'", cfg.BackendType)
	}
	return nil
}
