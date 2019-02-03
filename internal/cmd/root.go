package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultPort      = "8080"
	defaultInterface = "0.0.0.0"
	defaultLogLevel  = "INFO"
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
func Execute(info *ServiceInfo, serveHandler func(cmd *cobra.Command, args []string)) {
	rootCmd.Use = info.Name
	rootCmd.Short = info.Short
	rootCmd.Long = info.Long
	rootCmd.Version = info.Version

	for _, c := range rootCmd.Commands() {
		if c.Use == "serve" {
			c.Run = serveHandler
			break
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	rootCmd.PersistentFlags().String("logfile", "", "Log file with path for the service logging")
	rootCmd.PersistentFlags().String("loglevel", "", `Log level may be TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC. 
		Only will log lines of the equal or above severity`)

	viper.BindPFlag("logfile", rootCmd.PersistentFlags().Lookup("logfile"))
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
	viper.SetDefault("logfile", fmt.Sprintf("./%s.log", rootCmd.Use))
	viper.SetDefault("loglevel", defaultLogLevel)
}

func initConfig() {

	viper.AutomaticEnv()

	if cfgFlag := rootCmd.PersistentFlags().Lookup("config"); cfgFlag != nil {
		cfgFlag.Usage = fmt.Sprintf("config file (default is $HOME/.%s.yaml)", rootCmd.Use)
	}

	viper.SetDefault("logfile", fmt.Sprintf("./%s.log", rootCmd.Use))

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

		// Search config in home directory (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(fmt.Sprintf(".%s", rootCmd.Use))
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config file:", err)
	}
}
