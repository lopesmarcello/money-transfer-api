package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/lopesmarcello/money-transfer/internal/api"
	"github.com/lopesmarcello/money-transfer/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	fmt.Println("Setting env:")
	fmt.Println("User:", os.Getenv("MT_DATABASE_USER"))
	fmt.Println("Database:", os.Getenv("MT_DATABASE_NAME"))

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable search_path=public",
		os.Getenv("MT_DATABASE_USER"),
		os.Getenv("MT_DATABASE_PASSWORD"),
		os.Getenv("MT_DATABASE_HOST"),
		os.Getenv("MT_DATABASE_PORT"),
		os.Getenv("MT_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.API{
		Router:          chi.NewMux(),
		UserService:     services.NewUserService(pool),
		CurrencyService: services.NewCurrencyService(pool),
	}

	api.BindRoutes()
	fmt.Println("Starting server on port :8080")
	if err := http.ListenAndServe(":8080", api.Router); err != nil {
		panic(err)
	}
}
