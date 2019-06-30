package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// enableCmd represents the get command
var enableCmd = &cobra.Command{
	Use:          "enable",
	Short:        "Enable the vault-token-helper by (over)writing the ~/.vault config file",
	SilenceUsage: true, // Don't show help on error, just print the error

	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := homedir.Dir()
		if err != nil {
			return fmt.Errorf("Unable to determine home directory: %s", err)
		}
		path := filepath.Join(home, ".vault") // ~/.vault

		bin, err := filepath.Abs(os.Args[0])
		if err != nil {
			return fmt.Errorf("Unable to determine path to the vault-token-helper binary: %s", err)
		}

		// backup ~/.vault to ~/.vault.bak if it exists
		_, err = os.Stat(path)
		if !os.IsNotExist(err) {
			cmd.Println("Backing up existing ~/.vault file to ~/.vault.bak")
			bakPath := path + ".bak"
			if err := os.Rename(path, bakPath); err != nil {
				return err
			}
		}

		// write ~/.vault
		content := fmt.Sprintf("token_helper = \"%s\"", filepath.ToSlash(bin))
		return ioutil.WriteFile(path, []byte(content), 0640)
	},
}

func init() {
	RootCmd.AddCommand(enableCmd)
}
