# Сервис комментариев для новостного портала

Это полнофункциональный сервис для управления комментариями с использованием БД.

## Структура проекта

```
CommentService/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── comment_handler.go
│   ├── model/
│   │   └── comment.go
│   ├── repository/
│   │   └── comment_repository.go
│   ├── service/
│   │   └── comment_service.go
│   ├── moderator/
│   │   └── moderator.go
│   └── server/
│       └── server.go
├── migrations/
│   └── 001_create_comments.sql
├── docker-compose.yml
├── go.mod
└── README.md
```

## Запуск

1. Установите зависимости:
```bash
go mod tidy
```

2. Запустите базу данных с помощью Docker:
```bash
docker-compose up -d postgres
```

3. Дождитесь запуска базы данных и выполнения миграций

4. Запустите сервер:
```bash
go run cmd/server/main.go
```

Сервер запустится на `http://localhost:8081/`

## API Эндпоинты

- `POST /api/v1/comments` - создание комментария
- `GET /api/v1/comments/news/{news_id}` - комментарии к новости

## Примеры запросов

### Создать комментарий:
```bash
curl -X POST http://localhost:8081/api/v1/comments \
  -H "Content-Type: application/json" \
  -d '{
    "news_id": 123,
    "parent_id": null,
    "author": "Пользователь",
    "content": "Текст комментария"
  }'
```

### Получить комментарии к новости:
```bash
curl http://localhost:8081/api/v1/comments/news/123
```

## Асинхронная модерация

Сервис включает асинхронный процесс модерации комментариев, который запускается при старте приложения.
Модерация проверяет комментарии на наличие запрещенных слов: `qwerty`, `йцукен`, `zxcvbnm` (в любом регистре).
Комментарии с запрещенными словами отмечаются как неодобренные.