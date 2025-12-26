package service

import (
	"comment-service/internal/model"
	"comment-service/internal/repository"
	"errors"
	"time"
)

// CommentService - сервис для работы с комментариями
type CommentService struct {
	repo repository.CommentRepository
}

// NewCommentService создает новый сервис комментариев
func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

// CreateComment создает новый комментарий
func (s *CommentService) CreateComment(newsID int, parentID *int, author, content string) (*model.Comment, error) {
	// Валидация данных
	if author == "" {
		return nil, errors.New("автор обязателен")
	}

	if content == "" {
		return nil, errors.New("содержимое комментария обязательно")
	}

	if newsID <= 0 {
		return nil, errors.New("ID новости должен быть положительным числом")
	}

	// Создаем комментарий
	comment := &model.Comment{
		NewsID:    newsID,
		ParentID:  parentID,
		Author:    author,
		Content:   content,
		CreatedAt: time.Now(),
		Approved:  false, // по умолчанию не одобрен, ждет модерации
	}

	// Сохраняем в БД
	err := s.repo.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentsByNewsID возвращает комментарии для указанной новости
func (s *CommentService) GetCommentsByNewsID(newsID int) ([]model.Comment, error) {
	if newsID <= 0 {
		return nil, errors.New("ID новости должен быть положительным числом")
	}

	comments, err := s.repo.GetCommentsByNewsID(newsID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}