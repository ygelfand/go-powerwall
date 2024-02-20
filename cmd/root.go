package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type powerwallOptions struct {
	endpoint string
	password string
}

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
	viper.SetDefault("ENDPOINT", "https://192.168.91.1/tedapi")
	viper.BindEnv("PASSWORD")
	viper.BindEnv("ENDPOINT")
	o := &powerwallOptions{}
	rootCmd.PersistentFlags().StringVarP(&o.endpoint, "endpoint", "e", viper.GetString("ENDPOINT"), "powerwall endpoint url")
	rootCmd.PersistentFlags().StringVarP(&o.password, "password", "p", viper.GetString("PASSWORD"), "powerwall installer password")
	rootCmd.MarkPersistentFlagRequired("password")
	rootCmd.AddCommand(newProxyCmd(o))
	rootCmd.AddCommand(newDebugCmd(o))
	rootCmd.AddCommand(versionCmd)
	versionCmd.InheritedFlags().SetAnnotation("password", cobra.BashCompOneRequiredFlag, []string{"false"})
}