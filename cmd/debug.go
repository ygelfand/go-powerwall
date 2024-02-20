package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

func newDebugCmd(opts *powerwallOptions) *cobra.Command {
	debugCmd := &cobra.Command{
		Use:   "debug",
		Short: "run debug",
		Long:  `run some statuses for debug`,
		Run: func(cmd *cobra.Command, args []string) {
			pwr := powerwall.NewPowerwallGateway(opts.endpoint, opts.password)
			debug := pwr.RunQuery("DeviceControllerQuery", nil)
			var prettyJSON bytes.Buffer
			debug = pwr.RunQuery("ComponentsQuery", nil)
			err := json.Indent(&prettyJSON, []byte(*debug), "", "\t")
			if err != nil {
				fmt.Println("JSON parse error: ", err)
			}
			log.Println(string(prettyJSON.Bytes()))
			debug = pwr.GetConfig()
			err = json.Indent(&prettyJSON, []byte(*debug), "", "\t")
			if err != nil {
				fmt.Println("JSON parse error: ", err)
			}
			log.Println(string(prettyJSON.Bytes()))
		},
	}
	return debugCmd
}
