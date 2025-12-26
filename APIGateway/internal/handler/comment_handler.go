package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"api-gateway/internal/model"

	"github.com/gorilla/mux"
)

// CommentHandler - обработчик запросов комментариев
type CommentHandler struct {
	nextID int // простой счетчик ID для заглушки
}

// NewCommentHandler создает новый обработчик комментариев
func NewCommentHandler() *CommentHandler {
	return &CommentHandler{nextID: 1}
}

// AddComment - обработчик POST /api/v1/news/{id}/comments
func (h *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsIDStr := vars["id"]

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		http.Error(w, "ID новости должен быть числом", http.StatusBadRequest)
		return
	}

	// Декодируем тело запроса
	var req struct {
		Author   string `json:"author"`
		Content  string `json:"content"`
		ParentID *int   `json:"parent_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Валидация данных
	if req.Author == "" {
		http.Error(w, "Автор обязателен", http.StatusBadRequest)
		return
	}

	if req.Content == "" {
		http.Error(w, "Содержимое комментария обязательно", http.StatusBadRequest)
		return
	}

	// Создаем комментарий (заглушка)
	comment := model.Comment{
		ID:        h.nextID,
		NewsID:    newsID,
		ParentID:  req.ParentID,
		Author:    req.Author,
		Content:   req.Content,
		CreatedAt: time.Now(),
		Approved:  false, // по умолчанию не одобрен
	}

	h.nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}