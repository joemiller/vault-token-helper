package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "List tokens",
	SilenceUsage: true, // Don't show help on error, just print the error

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initBackend(); err != nil {
			return err
		}

		// tokens, err := backend.ListNew()
		// if err != nil {
		// 	return errors.Wrap(err, "Failed to read tokens from backend storage")
		// }
		// fmt.Println("VAULT_ADDR                               Created             LastModified")
		// fmt.Println("---------------------------------------- ------------------- -------------------")
		// // for _, t := range tokens {
		// // 	fmt.Println(t)
		// // }
		// for _, t := range tokens {
		// 	fmt.Printf("%-40s %-19s %-19s\n",
		// 		t.VaultAddr,
		// 		t.Created().Format("2006-01-02 15:04:05"),
		// 		t.LastModified().Format("2006-01-02 15:04:05"),
		// 	)
		// }

		keys, err := backend.List()
		if err != nil {
			return errors.Wrap(err, "Failed to read tokens from backend storage")
		}
		for _, t := range keys {
			cmd.Println(t)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
