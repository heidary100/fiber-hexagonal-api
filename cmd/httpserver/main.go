package httpserver

import (
	"context"
	"fmt"
	"log"
	"time"

	usersservice "github.com/heidary100/fiber-hexagonal-api/internal/core/service/user"
	usersrepository "github.com/heidary100/fiber-hexagonal-api/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/heidary100/fiber-hexagonal-api/internal/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RunFiberApp(port string) {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	userCollection := db.Collection("users")
	userRepo := usersrepository.NewRepo(userCollection)
	userService := usersservice.NewService(userRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the fiber-hexagonal-api mongo user manager!"))
	})
	api := app.Group("/api")
	routes.UserRouter(api, userService)
	defer cancel()
	log.Fatal(app.Listen(":" + port))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://root:rootpassword@localhost:27017").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("fiber-api")
	return db, cancel, nil
}
