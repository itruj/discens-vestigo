package main

import (
	"itruj/discens-vestigo/api"
	"itruj/discens-vestigo/api/config"
	database "itruj/discens-vestigo/api/store/db"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// make tailwind-build
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := config.MustLoadConfig()

	db := database.MustOpen(cfg.DatabaseName)

	srv := api.NewServer(db)
	http.ListenAndServe(":8080", srv.Routes())

	logger.Info("Server started", slog.String("dbName", cfg.DatabaseName))
}
