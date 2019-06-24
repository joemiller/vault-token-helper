package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:          "get",
	Short:        "Get the stored token for $VAULT_ADDR",
	Hidden:       true, // don't show in help output. This command is intended for Vault to invoke
	SilenceUsage: true, // Don't show help on error, just print the error

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initBackend(); err != nil {
			return err
		}
		vaultAddr := os.Getenv("VAULT_ADDR")
		if vaultAddr == "" {
			return errors.New("Missing VAULT_ADDR environment variable")
		}

		token, err := backend.Get(vaultAddr)
		if err != nil {
			return err
		}
		if token.Token != "" {
			fmt.Fprint(os.Stdout, token.Token) // no trailing newline when outputting token for Vault
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
