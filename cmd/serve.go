/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/heidary100/fiber-hexagonal-api/cmd/httpserver"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve command will run fiber api",
	Long:  `A longer description, Serve command will run fiber ap`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		mongoUri, _ := cmd.Flags().GetString("mongo-uri")
		dbName, _ := cmd.Flags().GetString("db-name")

		if port == "" {
			port = viper.Get("PORT").(string)
		}

		if mongoUri == "" {
			mongoUri = viper.Get("MONGO_URI").(string)
		}

		if dbName == "" {
			dbName = viper.Get("DB_NAME").(string)
		}

		httpserver.RunFiberApp(port, mongoUri, dbName)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "", "Port of fiber api")
	serveCmd.Flags().StringP("mongo-uri", "m", "", "MongoDB URI")
	serveCmd.Flags().StringP("db-name", "d", "", "MongoDB database name")
}
