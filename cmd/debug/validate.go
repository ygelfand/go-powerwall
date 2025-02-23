package debug

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ygelfand/go-powerwall/cmd/options"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
	"github.com/ygelfand/go-powerwall/internal/powerwall/queries"
)

// queryCmd represents the query command
func NewDebugValidateCmd(opts *options.PowerwallOptions) *cobra.Command {
	validateCmd := &cobra.Command{
		Use:       "validate",
		Short:     "run all saved queries",
		Long:      `Runs aall available queries for debug.`,
		ValidArgs: queries.QueryList(),
		Run: func(cmd *cobra.Command, args []string) {
			pwr := powerwall.NewPowerwallGateway(opts.Endpoint, opts.Password)
			var prettyJSON bytes.Buffer
			for _, q := range queries.QueryList() {
				fmt.Println(q)
				debug := pwr.RunQuery(q, nil)
				err := json.Indent(&prettyJSON, []byte(*debug), "", "\t")
				if err != nil {
					fmt.Println("JSON parse error: ", err)
				}
				fmt.Println(string(prettyJSON.Bytes()))
			}
		},
	}
	return validateCmd
}
