package main

import (
	"log"
	"net/http"

	"database/sql"
	"os"

	"github.com/Serega-D/bootdev-learning-chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // 1. Сначала читаем все нужные переменные из .env
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET is not set in .env")
    }

    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        log.Fatal("DB_URL is not set in .env")
    }

    platform := os.Getenv("PLATFORM")

    polkaSecret := os.Getenv("POLKA_KEY")

    // 2. Подключаемся к базе
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 3. Создаем ОДНУ структуру apiConfig и кладем туда ВСЁ
    apiCfg := &apiConfig{
        DB:        database.New(db),
        Platform:  platform,
        jwtSecret: jwtSecret, // Теперь секрет доступен в хендлерах через apiCfg
        polkaKey: polkaSecret,
    }

	const port = "8080"

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}
	log.Println("Successfully connected to database!")

	mux := http.NewServeMux()
	apiCfg.registerRoutes(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	server.ListenAndServe()
}
