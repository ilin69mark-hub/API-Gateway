package main

import (
	"log"
	"time"

	"comment-service/internal/server"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Подключаемся к базе данных
	// В реальном приложении строку подключения нужно получать из переменных окружения
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=comments_db sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Устанавливаем параметры подключения
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Создаем сервер
	srv := server.NewServer(db)

	// Запускаем сервер на порту 8081
	if err := srv.Run("8081"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}