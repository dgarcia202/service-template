package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts API service",
	Long:  `Stars the service in a given interface/port and processes requests`,
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
