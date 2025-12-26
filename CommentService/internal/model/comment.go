package model

import "time"

// Comment - комментарий
type Comment struct {
	ID         int       `json:"id" db:"id"`
	NewsID     int       `json:"news_id" db:"news_id"`
	ParentID   *int      `json:"parent_id,omitempty" db:"parent_id"` // nil для корневых комментариев
	Author     string    `json:"author" db:"author"`
	Content    string    `json:"content" db:"content"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Approved   bool      `json:"approved" db:"approved"` // true если прошел модерацию
	ModeratedAt *time.Time `json:"-" db:"moderated_at"`  // время модерации (внутреннее поле)
}