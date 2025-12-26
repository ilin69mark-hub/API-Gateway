package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"api-gateway/internal/model"

	"github.com/gorilla/mux"
)

// NewsHandler - обработчик запросов новостей
type NewsHandler struct{}

// NewNewsHandler создает новый обработчик новостей
func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

// GetNewsList - обработчик GET /api/v1/news
func (h *NewsHandler) GetNewsList(w http.ResponseWriter, r *http.Request) {
	// Возвращаем статичный список новостей (заглушка)
	news := []model.NewsShortDetailed{
		{
			ID:        1,
			Title:     "Новость 1",
			Published: time.Now().Add(-24 * time.Hour),
			Author:    "Иван Иванов",
		},
		{
			ID:        2,
			Title:     "Новость 2",
			Published: time.Now().Add(-12 * time.Hour),
			Author:    "Петр Петров",
		},
		{
			ID:        3,
			Title:     "Новость 3",
			Published: time.Now().Add(-6 * time.Hour),
			Author:    "Сидор Сидоров",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}

// GetNewsFilter - обработчик GET /api/v1/news/filter
func (h *NewsHandler) GetNewsFilter(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметры фильтрации
	author := r.URL.Query().Get("author")
	tag := r.URL.Query().Get("tag")
	fromDate := r.URL.Query().Get("from_date")
	toDate := r.URL.Query().Get("to_date")

	// Логируем параметры для отладки
	_ = author
	_ = tag
	_ = fromDate
	_ = toDate

	// Возвращаем тот же статичный список новостей (заглушка)
	news := []model.NewsShortDetailed{
		{
			ID:        1,
			Title:     "Новость 1",
			Published: time.Now().Add(-24 * time.Hour),
			Author:    "Иван Иванов",
		},
		{
			ID:        2,
			Title:     "Новость 2",
			Published: time.Now().Add(-12 * time.Hour),
			Author:    "Петр Петров",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}

// GetNewsDetail - обработчик GET /api/v1/news/{id}
func (h *NewsHandler) GetNewsDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID должен быть числом", http.StatusBadRequest)
		return
	}

	// Заглушка: возвращаем данные только для ID=1
	if id != 1 {
		http.Error(w, "Новость не найдена", http.StatusNotFound)
		return
	}

	news := model.NewsFullDetailed{
		NewsShortDetailed: model.NewsShortDetailed{
			ID:        1,
			Title:     "Полная новость 1",
			Published: time.Now().Add(-24 * time.Hour),
			Author:    "Иван Иванов",
		},
		Content: "Это полный текст новости 1. Здесь содержится подробная информация о событии.",
		Tags:    []string{"технологии", "новости", "IT"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}