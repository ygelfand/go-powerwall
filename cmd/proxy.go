package cmd

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/ygelfand/go-powerwall/cmd/options"
	"github.com/ygelfand/go-powerwall/internal/api"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

func newProxyCmd(opts *options.PowerwallOptions) *cobra.Command {
	o := &options.ProxyOptions{PowerwallOptions: opts}
	proxyCmd := &cobra.Command{
		Use:   "proxy",
		Short: "start powerwall proxy",
		Long:  `Start powerwall proxy server`,
		Run: func(cmd *cobra.Command, args []string) {
			pwr := powerwall.NewPowerwallGateway(o.Endpoint, o.Password)
			if !o.DebugMode {
				gin.SetMode(gin.ReleaseMode)
			}
			gin.ForceConsoleColor()
			app := api.NewApi(pwr, o.OnDemand)
			if !o.OnDemand {
				go pwr.PeriodicRefresh(time.Duration(o.RefreshInterval) * time.Second)
			}
			app.Run(o.ListenOn)
		},
	}
	proxyCmd.Flags().BoolVarP(&o.OnDemand, "ondemand", "o", false, "disable periodic refresh")
	proxyCmd.Flags().BoolVarP(&o.OnDemand, "full", "f", true, "start full authenticated portal proxy")
	proxyCmd.Flags().StringVarP(&o.ListenOn, "listen", "l", ":8080", "host:port to listen on")
	proxyCmd.Flags().Uint32VarP(&o.RefreshInterval, "refresh", "r", 30, "periodic refresh frequency in seconds")
	return proxyCmd
}
