package main

import (
	"log"

	"api-gateway/internal/server"
)

func main() {
	// Создаем сервер
	srv := server.NewServer()

	// Запускаем сервер на порту 8080
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}