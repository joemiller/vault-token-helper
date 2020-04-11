package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:          "erase",
	Short:        "(For use by vault) Erase the stored token for the current $VAULT_ADDR",
	SilenceUsage: true, // Don't show help on error, just print the error

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initBackend(); err != nil {
			return err
		}
		vaultAddr := os.Getenv("VAULT_ADDR")
		if vaultAddr == "" {
			return errors.New("Missing VAULT_ADDR environment variable")
		}

		vaultNamespace := os.Getenv("VAULT_NAMESPACE")
		if vaultNamespace != "" {
			vaultAddr += "/"
			vaultAddr += vaultNamespace
		}

		if err := backend.Erase(vaultAddr); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(eraseCmd)
}
