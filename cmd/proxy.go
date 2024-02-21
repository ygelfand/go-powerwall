package cmd

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/ygelfand/go-powerwall/internal/api"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

type proxyOptions struct {
	*powerwallOptions
	refreshInterval uint32
	onDemand        bool
	listenOn        string
}

func newProxyCmd(opts *powerwallOptions) *cobra.Command {
	o := &proxyOptions{powerwallOptions: opts}
	proxyCmd := &cobra.Command{
		Use:   "proxy",
		Short: "start powerwall proxy",
		Long:  `Start powerwall proxy server`,
		Run: func(cmd *cobra.Command, args []string) {
			pwr := powerwall.NewPowerwallGateway(o.endpoint, o.password)
			if !o.debugMode {
				gin.SetMode(gin.ReleaseMode)
			}
			gin.ForceConsoleColor()
			app := api.NewApi(pwr, o.onDemand)
			if !o.onDemand {
				go pwr.PeriodicRefresh(time.Duration(o.refreshInterval) * time.Second)
			}
			app.Run(o.listenOn)
		},
	}
	proxyCmd.Flags().BoolVarP(&o.onDemand, "ondemand", "o", false, "disable periodic refresh")
	proxyCmd.Flags().StringVarP(&o.listenOn, "listen", "l", ":8080", "host:port to listen on")
	proxyCmd.Flags().Uint32VarP(&o.refreshInterval, "refresh", "r", 30, "periodic refresh frequency in seconds")
	return proxyCmd
}
