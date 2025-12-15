package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	db "go-user-api/db/sqlc"
	"go-user-api/internal/handler"
	"go-user-api/internal/logger"
	"go-user-api/internal/middleware"
	"go-user-api/internal/repository"
	"go-user-api/internal/routes"
	"go-user-api/internal/service"
)

func main() {
	// Init zap logger
	logger.Init()
	defer logger.Sync()

	logger.Log.Info("starting server")

	dsn := "postgres://postgres:Anagha%402906@localhost:5432/userdb?sslmode=disable"

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Log.Fatal("failed to open database", zap.Error(err))
	}

	if err := conn.Ping(); err != nil {
		logger.Log.Fatal("database connection failed", zap.Error(err))
	}

	logger.Log.Info("database connected")

	queries := db.New(conn)
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	app := fiber.New()

	// request logging middleware
	app.Use(middleware.RequestLogger())

	routes.Register(app, h)

	logger.Log.Info("server listening on :8080")
	logger.Log.Fatal("server stopped", zap.Error(app.Listen(":8080")))
}
