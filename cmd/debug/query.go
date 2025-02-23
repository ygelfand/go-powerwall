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
func NewDebugQueryCmd(opts *options.PowerwallOptions) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:       "query [queryName]",
		Short:     "run a saved query",
		Long:      `Runs an available query for debug.`,
		ValidArgs: queries.QueryList(),
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			pwr := powerwall.NewPowerwallGateway(opts.Endpoint, opts.Password)
			debug := pwr.RunQuery(args[0], nil)
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(*debug), "", "\t")
			if err != nil {
				fmt.Println("JSON parse error: ", err)
			}

			fmt.Println(string(prettyJSON.Bytes()))
		},
	}
	originalUsageFunc := queryCmd.UsageFunc()

	queryCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		originalUsageFunc(cmd)
		fmt.Fprintf(cmd.OutOrStderr(), "\nKnown queries:\n")
		for _, arg := range cmd.ValidArgs {
			fmt.Fprintf(cmd.OutOrStderr(), "  %s\n", arg)
		}
		fmt.Fprintf(cmd.OutOrStderr(), "\n")
		return nil
	})
	return queryCmd
}
