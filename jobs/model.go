package jobs

import "time"

type Job struct {
	ID      string `json:"id"`      // UUID (veya snowflake)
	URL     string `json:"url"`     // Gönderilecek endpoint
	Method  string `json:"method"`  // POST, GET, PUT, DELETE
	Headers string `json:"headers"` // JSON string olarak saklanabilir (map değil!)
	Body    string `json:"body"`    // Raw JSON / text body

	ExecuteAt  time.Time `json:"execute_at"`  // Ne zaman çalışacak
	Status     string    `json:"status"`      // pending, running, done, failed, cancelled
	RetryCount int       `json:"retry_count"` // Kaç kez denendi
	MaxRetries int       `json:"max_retries"` // Kaç kereye kadar retry edilsin

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TokenID string `json:"token_id"` // Bu job hangi API token ile ilişkili
}
