package models

import (
	"encoding/json"
	"time"
)

type Job struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
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

// UnmarshalJSON var ise otomatik olarak çözümlenir.
func (j *Job) UnmarshalJSON(data []byte) error {
	type Alias Job
	aux := &struct {
		ExecuteAt string `json:"execute_at"`
		*Alias
	}{
		Alias: (*Alias)(j),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Tarih formatını çözümle
	if aux.ExecuteAt == "" {
		return nil
	}

	loc, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation("2006-01-02 15:04", aux.ExecuteAt, loc)
	if err != nil {
		return err
	}
	j.ExecuteAt = t

	return nil
}
