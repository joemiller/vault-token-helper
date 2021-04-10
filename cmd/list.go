package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sync"
	"text/tabwriter"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().BoolP("extended", "e", false, "Lookup and print extended details about each token by querying each Vault server")
	RootCmd.AddCommand(listCmd)
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "List tokens",
	SilenceUsage: true, // Don't show help on error, just print the error

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initBackend(); err != nil {
			return err
		}

		extended := cmd.Flag("extended").Value.String() == "true"
		if extended {
			if err := extendedList(); err != nil {
				return err
			}
			return nil
		}

		// simple list, just print vault-addrs
		addrs, err := backend.List()
		if err != nil {
			return errors.Wrap(err, "Failed to read tokens from backend storage")
		}
		for _, addr := range addrs {
			cmd.Println(addr)
		}

		return nil
	},
}

func extendedList() error {
	wg := &sync.WaitGroup{}

	addrs, err := backend.List()
	if err != nil {
		return errors.Wrap(err, "Failed to read tokens from backend storage")
	}

	// fan-out parallel token lookups, fan-in to result channels. Output and Error
	// channels are rendered separately to ensure correct aligntment in the outputted table
	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 5, 0, 2, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintln(w, "VAULT_ADDR\tdisplay_name\tttl\trenewable\tpolicies\t")
	fmt.Fprintln(w, "----------\t------------\t---\t---------\t--------\t")

	outCh := make(chan string, len(addrs))
	errCh := make(chan error, len(addrs))

	for _, addr := range addrs {
		wg.Add(1)

		go func(addr string) {
			defer wg.Done()

			// Fetch the stored token from the backend
			token, err := backend.Get(addr)
			if err != nil {
				errCh <- fmt.Errorf("%s\t** ERROR **\t%s", addr, err)
				return
			}

			// lookup the token in the vault instance at 'addr'
			vcfg := vault.DefaultConfig() // VAULT_ env vars
			vcfg.Timeout = 5 * time.Second
			vcfg.MaxRetries = 1

			// XXX: We store the VAULT_ADDR + VAULT_NAMESPACE in the credential store as a single
			// string, eg:
			//
			//   VAULT_ADDR=https://vault:8200  VAULT_NAMESPACE=foo is stored as "https://vault:8200/foo"
			//
			// But this is not a valid VAULT_ADDR. To workaround this we parse the string and assume
			// a Path element is the VAULT_NAMESPACE.
			parsedURL, err := url.Parse(addr)
			if err != nil {
				errCh <- fmt.Errorf("%s\t** ERROR **\t%s", addr, err)
				return
			}
			vcfg.Address = fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
			namespace := parsedURL.Path

			client, err := vault.NewClient(vcfg)
			if err != nil {
				errCh <- fmt.Errorf("%s\t** ERROR **\t%s", addr, err)
				return
			}
			client.SetNamespace(namespace)

			client.SetToken(token.Token)
			s, err := client.Auth().Token().LookupSelf()
			if err != nil {
				errCh <- fmt.Errorf("%s\t** ERROR **\t%s", addr, err)
				return
			}

			ttl, err := s.Data["ttl"].(json.Number).Int64()
			if err != nil {
				errCh <- fmt.Errorf("%s\t** ERROR **\t%s", addr, err)
				return
			}
			outCh <- fmt.Sprintf("%s\t%s\t%s\t%t\t%s",
				addr,
				s.Data["display_name"],
				time.Duration(ttl)*time.Second,
				s.Data["renewable"],
				s.Data["policies"],
			)
		}(addr)
	}
	// close the result channels when all lookups are complete
	go func() {
		wg.Wait()
		close(outCh)
		close(errCh)
	}()

	for i := range outCh {
		fmt.Fprintln(w, i)
	}
	for i := range errCh {
		fmt.Fprintln(w, i)
	}
	w.Flush()
	return nil
}
