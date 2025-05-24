package jobs

import (
	"net/http"
	"strconv"

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
