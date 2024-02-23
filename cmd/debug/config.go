package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ygelfand/go-powerwall/cmd/options"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

func NewDebugConfigCmd(opts *options.PowerwallOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "config [queryName]",
		Short: "pull config",
		Long:  `Pulls json config from firewall.`,
		Run: func(cmd *cobra.Command, args []string) {
			pwr := powerwall.NewPowerwallGateway(opts.Endpoint, opts.Password)
			debug := pwr.GetConfig()
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(*debug), "", "\t")
			if err != nil {
				fmt.Println("JSON parse error: ", err)
			}
			log.Println(string(prettyJSON.Bytes()))
		},
	}
}
