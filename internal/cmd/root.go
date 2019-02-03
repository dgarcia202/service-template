package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServiceInfo holds basic service info for the root command
type ServiceInfo struct {
	Name, Short, Long, Version string
}

var cfgFile string

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute runs the root command logic
func Execute(info *ServiceInfo) {
	rootCmd.Use = info.Name
	rootCmd.Short = info.Short
	rootCmd.Long = info.Long
	rootCmd.Version = info.Version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
}

func initConfig() {

	if cfgFlag := rootCmd.PersistentFlags().Lookup("config"); cfgFlag != nil {
		cfgFlag.Usage = fmt.Sprintf("config file (default is $HOME/.%s.yaml)", rootCmd.Use)
	}

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// TODO: deal with the dynamic config file name

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(fmt.Sprintf(".%s.yaml", rootCmd.Use))
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config file:", err)
	}
}
