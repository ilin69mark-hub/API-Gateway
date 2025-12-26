-- Миграция для создания таблицы комментариев

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    news_id INTEGER NOT NULL,
    parent_id INTEGER,
    author VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    approved BOOLEAN DEFAULT FALSE,
    moderated_at TIMESTAMP
);

-- Создаем индексы для улучшения производительности
CREATE INDEX idx_comments_news_id ON comments (news_id);
CREATE INDEX idx_comments_parent_id ON comments (parent_id);
CREATE INDEX idx_comments_created_at ON comments (created_at);
CREATE INDEX idx_comments_approved ON comments (approved);