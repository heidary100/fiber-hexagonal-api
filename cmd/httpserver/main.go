package httpserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	moviesservice "github.com/heidary100/fiber-hexagonal-api/internal/core/service/movie"
	usersservice "github.com/heidary100/fiber-hexagonal-api/internal/core/service/user"
	moviesrepository "github.com/heidary100/fiber-hexagonal-api/internal/repositories/movie"
	usersrepository "github.com/heidary100/fiber-hexagonal-api/internal/repositories/user"
	"github.com/heidary100/fiber-hexagonal-api/internal/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RunFiberApp(port string, mongoUri string, dbName string) {
	fmt.Println(port, mongoUri, dbName)
	db, cancel, err := databaseConnection(mongoUri, dbName)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	userCollection := db.Collection("users")
	userRepo := usersrepository.NewRepo(userCollection)
	userService := usersservice.NewService(userRepo)

	movieCollection := db.Collection("movies")
	movieRepo := moviesrepository.NewRepo(movieCollection)
	movieService := moviesservice.NewService(movieRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the fiber-hexagonal-api mongo user manager!"))
	})
	api := app.Group("/api")
	routes.UserRouter(api, userService)
	routes.MovieRouter(api, movieService)
	defer cancel()
	log.Fatal(app.Listen(":" + port))
}

func databaseConnection(mongoUri string, dbName string) (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(dbName)
	return db, cancel, nil
}
