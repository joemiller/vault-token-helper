package cmd

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/joemiller/vault-token-helper/pkg/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:          "store",
	Short:        "(For use by vault) Store a token (from stdin) for the current $VAULT_ADDR",
	SilenceUsage: true, // Don't show help on error, just print the error

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initBackend(); err != nil {
			return err
		}
		vaultAddr := os.Getenv("VAULT_ADDR")
		if vaultAddr == "" {
			return errors.New("Missing VAULT_ADDR environment variable")
		}

		stdin, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return errors.Wrap(err, "Failed to read token from STDIN")
		}

		token := store.Token{
			VaultAddr: vaultAddr,
			Token:     strings.TrimSuffix(string(stdin), "\n"),
		}
		if err := backend.Store(token); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(storeCmd)
}
