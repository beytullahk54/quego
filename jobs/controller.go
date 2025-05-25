package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{DB: db}
}

func (c *Controller) GetJobs(ctx *gin.Context) {
	var properties []Job
	result := c.DB.Find(&properties)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, properties)
}

func (c *Controller) GetJobByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var property Job
	result := c.DB.First(&property, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, property)
}

func (c *Controller) CreateJob(ctx *gin.Context) {
	var newProperty Job
	if err := ctx.ShouldBindJSON(&newProperty); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := c.DB.Create(&newProperty)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newProperty)
}

func (c *Controller) UpdateJob(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProperty Job
	if err := ctx.ShouldBindJSON(&updatedProperty); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingProperty Job
	if err := c.DB.First(&existingProperty, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	result := c.DB.Model(&existingProperty).Updates(Job{
		URL:        updatedProperty.URL,
		Method:     updatedProperty.Method,
		Headers:    updatedProperty.Headers,
		Body:       updatedProperty.Body,
		ExecuteAt:  updatedProperty.ExecuteAt,
		Status:     updatedProperty.Status,
		RetryCount: updatedProperty.RetryCount,
		MaxRetries: updatedProperty.MaxRetries,
		TokenID:    updatedProperty.TokenID,
	})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	updatedProperty.ID = strconv.Itoa(id)
	ctx.JSON(http.StatusOK, updatedProperty)
}

func (c *Controller) DeleteJob(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := c.DB.Delete(&Job{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// CheckAndProcessJobs, jobs tablosunu kontrol eder ve gerekli istekleri atar
func (c *Controller) CheckAndProcessJobs() {
	var jobs []Job
	now := time.Now()

	// Şu anki zamandan önce çalışması gereken işleri al
	err := c.DB.Where("execute_at <= ? AND status = ?", now, "1").Find(&jobs).Error
	if err != nil {
		log.Printf("Jobs kontrol edilirken hata oluştu: %v", err)
		return
	}

	for _, job := range jobs {
		// İsteği gönder
		err := sendRequest(job, c)
		if err != nil {
			log.Printf("İstek gönderilirken hata oluştu (Job ID: %d): %v", job.ID, err)
			continue
		}

		// Bir sonraki çalışma zamanını güncelle (5 dakika sonra)
		/*job.NextRun = now.Add(5 * time.Minute)
		if err := c.DB.Save(&job).Error; err != nil {
			log.Printf("Job güncellenirken hata oluştu (Job ID: %d): %v", job.ID, err)
		}*/
	}
}

// sendRequest, belirtilen job için HTTP isteği gönderir
func sendRequest(job Job, c *Controller) error {
	// İstek gövdesini hazırla
	requestBody, err := json.Marshal(map[string]interface{}{
		"job_id":      job.ID,
		"url":         job.URL,
		"method":      job.Method,
		"headers":     job.Headers,
		"body":        job.Body,
		"execute_at":  job.ExecuteAt,
		"status":      job.Status,
		"retry_count": job.RetryCount,
		"max_retries": job.MaxRetries,
		"token_id":    job.TokenID,
	})
	if err != nil {
		return err
	}

	log.Println("İstek gönderiliyor:", job.URL)
	// HTTP isteği gönder
	req, err := http.NewRequest(job.Method, job.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// HTTP isteğini gönder
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("İstek gönderildi:", job.ID)

	var existingProperty Job
	if err := c.DB.First(&existingProperty, job.ID).Error; err != nil {
		return err
	}

	result := c.DB.Model(&existingProperty).Updates(Job{
		Status: "2",
	})

	if result.Error != nil {
		return result.Error
	}

	// Yanıt durumunu kontrol et
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("beklenmeyen yanıt kodu: %d", resp.StatusCode)
	}

	return nil
}
