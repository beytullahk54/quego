package users

import (
	"time"

	"github.com/go-playground/validator/v10"
	"quego.com/gin-crud/internal/models"
)

// JobDTO API isteklerinde kullanılan veri transferi için
// validation kuralları burada tanımlanıyor
type UserDTO struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

// Validate DTO'nun validation kurallarını kontrol eder
func (j *UserDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(j)
}

// ToJob DTO'yu Job struct'ine dönüştürür
func (j *UserDTO) ToJob() *models.User {
	job := &models.User{
		Name:     j.Name,
		Email:    j.Email,
		Password: j.Password,
	}

	// CreatedAt ve UpdatedAt'ı parse et
	if j.CreatedAt != "" {
		loc, _ := time.LoadLocation("Europe/Istanbul")
		t, _ := time.ParseInLocation("2006-01-02 15:04", j.CreatedAt, loc)
		job.CreatedAt = t
	}

	return job
}
