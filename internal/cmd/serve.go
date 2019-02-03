package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultPort      = "8080"
	defaultInterface = "0.0.0.0"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts API service",
	Long:  `Stars the service in a given interface/port and processes requests`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()

		r.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		r.Run(fmt.Sprintf("%s:%s", viper.GetString("address"), viper.GetString("port")))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", defaultPort, "Port where to listen for request")
	serveCmd.Flags().StringP("address", "a", defaultInterface, "Address where to listen for request")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("address", serveCmd.Flags().Lookup("address"))
	viper.SetDefault("port", defaultPort)
	viper.SetDefault("address", defaultInterface)
}
