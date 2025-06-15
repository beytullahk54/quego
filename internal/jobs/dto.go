package jobs

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// JobDTO API isteklerinde kullanılan veri transferi için
// validation kuralları burada tanımlanıyor
type JobDTO struct {
	URL         string    `json:"url" validate:"required,url"`
	Method      string    `json:"method" validate:"required,oneof=GET POST PUT DELETE"`
	Headers     string    `json:"headers"`
	Body        string    `json:"body"`
	ExecuteAt   string    `json:"execute_at" validate:"required,datetime=2006-01-02 15:04"`
	Status      string    `json:"status"`
	RetryCount  int       `json:"retry_count"`
	MaxRetries  int       `json:"max_retries"`
	TokenID     string    `json:"token_id"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

// Validate DTO'nun validation kurallarını kontrol eder
func (j *JobDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(j)
}

// ToJob DTO'yu Job struct'ine dönüştürür
func (j *JobDTO) ToJob() *Job {
	job := &Job{
		URL:        j.URL,
		Method:     j.Method,
		Headers:    j.Headers,
		Body:       j.Body,
		Status:     j.Status,
		RetryCount: j.RetryCount,
		MaxRetries: j.MaxRetries,
		TokenID:    j.TokenID,
	}

	// ExecuteAt'i parse et
	if j.ExecuteAt != "" {
		loc, _ := time.LoadLocation("Europe/Istanbul")
		t, _ := time.ParseInLocation("2006-01-02 15:04", j.ExecuteAt, loc)
		job.ExecuteAt = t
	}

	// CreatedAt ve UpdatedAt'ı parse et
	if j.CreatedAt != "" {
		loc, _ := time.LoadLocation("Europe/Istanbul")
		t, _ := time.ParseInLocation("2006-01-02 15:04", j.CreatedAt, loc)
		job.CreatedAt = t
	}

	if j.UpdatedAt != "" {
		loc, _ := time.LoadLocation("Europe/Istanbul")
		t, _ := time.ParseInLocation("2006-01-02 15:04", j.UpdatedAt, loc)
		job.UpdatedAt = t
	}

	return job
}
