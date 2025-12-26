# API Gateway для новостного портала

Это API Gateway для новостного портала с заглушками для интеграции с фронтендом.

## Структура проекта

```
APIGateway/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── news_handler.go
│   │   └── comment_handler.go
│   ├── model/
│   │   ├── news.go
│   │   └── comment.go
│   └── server/
│       └── server.go
├── go.mod
└── README.md
```

## Запуск

1. Установите зависимости:
```bash
go mod tidy
```

2. Запустите сервер:
```bash
go run cmd/server/main.go
```

Сервер запустится на `http://localhost:8080/`

## API Эндпоинты

### Новости

- `GET /api/v1/news` - список новостей
- `GET /api/v1/news/filter` - фильтр новостей (параметры: author, tag, from_date, to_date)
- `GET /api/v1/news/{id}` - детальная новость

### Комментарии

- `POST /api/v1/news/{id}/comments` - добавление комментария к новости

## Примеры запросов

### Получить список новостей:
```bash
curl http://localhost:8080/api/v1/news
```

### Получить новость по ID:
```bash
curl http://localhost:8080/api/v1/news/1
```

### Добавить комментарий:
```bash
curl -X POST http://localhost:8080/api/v1/news/1/comments \
  -H "Content-Type: application/json" \
  -d '{"author": "Иван", "content": "Отличная новость!"}'
```