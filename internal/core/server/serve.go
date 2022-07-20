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
			//repositories
			userRepository := repositories.NewUserRepository()
			filmRepository := repositories.NewFilmRepository()
			musicRepository := repositories.NewMusicRepository()
			//services
			userService := services.NewUserService(userRepository)
			filmService := services.NewFilmService(filmRepository)
			musicService := services.NewMusicService(musicRepository)
			//handlers
			userHandlers := handlers.NewUserHandlers(userService)
			filmHandlers := handlers.NewFilmHandlers(filmService)
			musicHandlers := handlers.NewMusicHandlers(musicService)
			//server
			httpServer := NewServer(
				userHandlers,
				filmHandlers,
				musicHandlers,
			)
			return httpServer.Initialize(":" + viper.GetString("PORT"))
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}
