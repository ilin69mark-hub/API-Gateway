package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"comment-service/internal/handler"
	"comment-service/internal/moderator"
	"comment-service/internal/repository"
	"comment-service/internal/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Server - структура сервера сервиса комментариев
type Server struct {
	httpServer *http.Server
	db         *sqlx.DB
}

// NewServer создает новый сервер
func NewServer(db *sqlx.DB) *Server {
	return &Server{db: db}
}

// Run запускает HTTP-сервер на указанном порту
func (s *Server) Run(port string) error {
	// Создаем маршрутизатор
	router := mux.NewRouter()

	// Создаем зависимости
	commentRepo := repository.NewCommentRepository(s.db)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService)

	// Настраиваем маршруты
	api := router.PathPrefix("/api/v1").Subrouter()

	// Маршруты для комментариев
	api.HandleFunc("/comments", commentHandler.CreateComment).Methods("POST")
	api.HandleFunc("/comments/news/{news_id}", commentHandler.GetCommentsByNewsID).Methods("GET")

	// Создаем HTTP-сервер
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запускаем асинхронную модерацию комментариев
	go moderator.StartModerationQueue(commentRepo)

	// Логируем запуск сервера
	log.Printf("Сервер комментариев запущен на порту %s", port)

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
