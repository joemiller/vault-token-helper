package cmd

import (
	"github.com/joemiller/vault-token-helper/pkg/store"
	"github.com/spf13/cobra"
)

// backendsCmd represents the get command
var backendsCmd = &cobra.Command{
	Use:   "backends",
	Short: "List the available backends on the current platform",
	Long: `List the available backends on the current platform.

The backends are listed in the order of preference they will
be tried if no backend is selected in the configuration file.
`,

	Run: func(cmd *cobra.Command, args []string) {
		for _, i := range store.AvailableBackends() {
			cmd.Println(i) // TODO would be nice to use cmd.print, but need separate out channels to merge in cobra
			// fmt.Println(i)
		}
	},
}

func init() {
	RootCmd.AddCommand(backendsCmd)
}
