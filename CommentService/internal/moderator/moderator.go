package moderator

import (
	"comment-service/internal/repository"
	"log"
	"strings"
	"time"
)

// StartModerationQueue запускает асинхронный процесс модерации комментариев
func StartModerationQueue(repo repository.CommentRepository) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// 1. Получаем непромодерированные комментарии
		comments, err := repo.GetUnmoderatedComments()
		if err != nil {
			log.Printf("Ошибка при получении непромодерированных комментариев: %v", err)
			continue
		}

		// 2. Проверяем каждый комментарий на запрещенные слова и обновляем статус
		for _, comment := range comments {
			approved := isCommentApproved(comment.Content)
			
			// 3. Обновляем статус approved
			err := repo.UpdateCommentApproval(comment.ID, approved)
			if err != nil {
				log.Printf("Ошибка при обновлении статуса комментария %d: %v", comment.ID, err)
				continue
			}

			// 4. Логируем результат
			status := "одобрен"
			if !approved {
				status = "отклонен"
			}
			log.Printf("Комментарий %d %s модерацией", comment.ID, status)
		}
	}
}

// isCommentApproved проверяет, проходит ли комментарий модерацию
// Возвращает true, если комментарий одобрен, false - если нет
func isCommentApproved(content string) bool {
	// Запрещенные слова: qwerty, йцукен, zxcvbnm (в любом регистре)
	forbiddenWords := []string{"qwerty", "йцукен", "zxcvbnm"}

	contentLower := strings.ToLower(content)
	for _, word := range forbiddenWords {
		if strings.Contains(contentLower, word) {
			return false
		}
	}

	return true
}