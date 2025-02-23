package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ygelfand/go-powerwall/cmd/debug"
	_ "github.com/ygelfand/go-powerwall/cmd/debug"
	"github.com/ygelfand/go-powerwall/cmd/options"
)

func newDebugCmd(opts *options.PowerwallOptions) *cobra.Command {
	debugCmd := &cobra.Command{
		Use:   "debug",
		Short: "run debug",
		Long:  `run some statuses for debug`,
	}
	debugCmd.AddCommand(debug.NewDebugQueryCmd(opts))
	debugCmd.AddCommand(debug.NewDebugConfigCmd(opts))
	debugCmd.AddCommand(debug.NewDebugValidateCmd(opts))
	return debugCmd
}
