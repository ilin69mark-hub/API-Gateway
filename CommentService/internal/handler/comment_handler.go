package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"comment-service/internal/model"
	"comment-service/internal/service"

	"github.com/gorilla/mux"
)

// CommentHandler - обработчик запросов комментариев
type CommentHandler struct {
	service *service.CommentService
}

// NewCommentHandler создает новый обработчик комментариев
func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// CreateComment - обработчик POST /api/v1/comments
func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	// Декодируем тело запроса
	var req struct {
		NewsID   int    `json:"news_id"`
		ParentID *int   `json:"parent_id"`
		Author   string `json:"author"`
		Content  string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Создаем комментарий через сервис
	comment, err := h.service.CreateComment(req.NewsID, req.ParentID, req.Author, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// GetCommentsByNewsID - обработчик GET /api/v1/comments/news/{news_id}
func (h *CommentHandler) GetCommentsByNewsID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsIDStr := vars["news_id"]

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		http.Error(w, "ID новости должен быть числом", http.StatusBadRequest)
		return
	}

	comments, err := h.service.GetCommentsByNewsID(newsID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}