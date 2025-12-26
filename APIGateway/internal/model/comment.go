package model

import "time"

// Comment - комментарий
type Comment struct {
	ID        int       `json:"id"`
	NewsID    int       `json:"news_id"`
	ParentID  *int      `json:"parent_id,omitempty"` // nil для корневых комментариев
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Approved  bool      `json:"approved"` // true если прошел модерацию
}