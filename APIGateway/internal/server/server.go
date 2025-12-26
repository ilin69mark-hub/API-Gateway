package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-gateway/internal/handler"

	"github.com/gorilla/mux"
)

// Server - структура сервера API Gateway
type Server struct {
	httpServer *http.Server
}

// NewServer создает новый сервер
func NewServer() *Server {
	return &Server{}
}

// Run запускает HTTP-сервер на указанном порту
func (s *Server) Run(port string) error {
	// Создаем маршрутизатор
	router := mux.NewRouter()

	// Создаем обработчики
	newsHandler := handler.NewNewsHandler()
	commentHandler := handler.NewCommentHandler()

	// Настраиваем маршруты
	api := router.PathPrefix("/api/v1").Subrouter()

	// Маршруты для новостей
	api.HandleFunc("/news", newsHandler.GetNewsList).Methods("GET")
	api.HandleFunc("/news/filter", newsHandler.GetNewsFilter).Methods("GET")
	api.HandleFunc("/news/{id}", newsHandler.GetNewsDetail).Methods("GET")

	// Маршрут для добавления комментария к новости
	api.HandleFunc("/news/{id}/comments", commentHandler.AddComment).Methods("POST")

	// Создаем HTTP-сервер
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Логируем запуск сервера
	log.Printf("Сервер запущен на порту %s", port)

	// Запускаем сервер в горутине
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()

	// Ожидаем сигнал остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Ждем сигнал остановки
	<-stop
	log.Println("Получен сигнал остановки сервера...")

	// Создаем контекст для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Останавливаем сервер
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при остановке сервера: %v", err)
		return err
	}

	log.Println("Сервер остановлен")
	return nil
}