package repository

import (
	"comment-service/internal/model"

	"github.com/jmoiron/sqlx"
)

// CommentRepository - интерфейс для работы с комментариями
type CommentRepository interface {
	CreateComment(comment *model.Comment) error
	GetCommentsByNewsID(newsID int) ([]model.Comment, error)
	GetUnmoderatedComments() ([]model.Comment, error)
	UpdateCommentApproval(id int, approved bool) error
}

// commentRepository - реализация репозитория комментариев
type commentRepository struct {
	db *sqlx.DB
}

// NewCommentRepository создает новый репозиторий комментариев
func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepository{db: db}
}

// CreateComment создает новый комментарий
func (r *commentRepository) CreateComment(comment *model.Comment) error {
	query := `
		INSERT INTO comments (news_id, parent_id, author, content, created_at, approved)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err := r.db.QueryRow(
		query,
		comment.NewsID,
		comment.ParentID,
		comment.Author,
		comment.Content,
		comment.CreatedAt,
		comment.Approved,
	).Scan(&comment.ID)

	return err
}

// GetCommentsByNewsID возвращает все комментарии для указанной новости
func (r *commentRepository) GetCommentsByNewsID(newsID int) ([]model.Comment, error) {
	query := `
		SELECT id, news_id, parent_id, author, content, created_at, approved
		FROM comments
		WHERE news_id = $1 AND approved = true
		ORDER BY created_at ASC`

	var comments []model.Comment
	err := r.db.Select(&comments, query, newsID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// GetUnmoderatedComments возвращает непромодерированные комментарии
func (r *commentRepository) GetUnmoderatedComments() ([]model.Comment, error) {
	query := `
		SELECT id, news_id, parent_id, author, content, created_at, approved
		FROM comments
		WHERE approved = false`

	var comments []model.Comment
	err := r.db.Select(&comments, query)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// UpdateCommentApproval обновляет статус одобрения комментария
func (r *commentRepository) UpdateCommentApproval(id int, approved bool) error {
	query := `
		UPDATE comments
		SET approved = $1, moderated_at = NOW()
		WHERE id = $2`

	_, err := r.db.Exec(query, approved, id)
	return err
}