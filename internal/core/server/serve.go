package server

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/handlers"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/repositories"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start a HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			mongoConn := viper.GetString("MONGO_URI")
			//repositories
			userRepository := repositories.NewUserRepository(mongoConn)
			//services
			userService := services.NewUserService(userRepository)
			//handlers
			userHandlers := handlers.NewUserHandlers(userService)
			//server
			httpServer := NewServer(
				userHandlers,
			)
			return httpServer.Initialize(":" + viper.GetString("PORT"))
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}
