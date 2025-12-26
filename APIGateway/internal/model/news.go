package model

import "time"

// NewsShortDetailed - краткая информация о новости
type NewsShortDetailed struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Published time.Time `json:"published"`
	Author    string    `json:"author"`
}

// NewsFullDetailed - полная информация о новости
type NewsFullDetailed struct {
	NewsShortDetailed
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}