package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Golang CLI Template project",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is ./configs/.config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName(".config")
	}

	viper.SetConfigType("yaml")

	//viper.SetEnvPrefix("server")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to load config file:", viper.ConfigFileUsed())
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to load config file:", viper.ConfigFileUsed())
	}
}
