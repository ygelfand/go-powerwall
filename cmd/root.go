package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ygelfand/go-powerwall/cmd/options"
)

// set at build time
var debugMode = "true"

var rootCmd = &cobra.Command{
	Use:   "go-powerwall",
	Short: "go powerwall proxy",
	Long:  `Go powerwall proxy `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix("POWERWALL")
	viper.SetDefault("ENDPOINT", "https://192.168.91.1/")
	viper.BindEnv("PASSWORD")
	viper.BindEnv("ENDPOINT")
	o := &options.PowerwallOptions{DebugMode: debugMode == "true"}
	rootCmd.PersistentFlags().StringVarP(&o.Endpoint, "endpoint", "e", viper.GetString("ENDPOINT"), "powerwall endpoint url")
	rootCmd.PersistentFlags().StringVarP(&o.Password, "password", "p", viper.GetString("PASSWORD"), "powerwall installer password")
	rootCmd.MarkPersistentFlagRequired("password")
	rootCmd.AddCommand(newProxyCmd(o))
	rootCmd.AddCommand(newDebugCmd(o))
	rootCmd.AddCommand(versionCmd)
	versionCmd.InheritedFlags().SetAnnotation("password", cobra.BashCompOneRequiredFlag, []string{"false"})
}
