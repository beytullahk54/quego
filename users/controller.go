package users

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

func (c *Controller) GetUsers(ctx *gin.Context) {
	// Boş bir Book slice'ı (dizi) oluştur
	var users []User

	// Veritabanından tüm kitapları çek
	// c.DB: Controller'ın veritabanı bağlantısı
	// Find(): Tüm kayıtları getirir
	// &books: books slice'ının adresini ver (pointer)
	result := c.DB.Find(&users)

	// Hata kontrolü
	if result.Error != nil {
		// Hata varsa 500 (Internal Server Error) döndür
		// gin.H: JSON formatında veri oluşturur
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Başarılı ise 200 (OK) status kodu ile kitapları döndür
	ctx.JSON(http.StatusOK, users)
}

func (c *Controller) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user User
	result := c.DB.First(&user, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var newDTO UserDTO
	if err := ctx.ShouldBindJSON(&newDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation kontrolü
	if err := newDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Gerekli alanlar eksik veya hatalı",
			"details": err.Error(),
		})
		return
	}

	newUser := newDTO.ToJob()

	result := c.DB.Create(&newUser)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedDTO UserDTO
	if err := ctx.ShouldBindJSON(&updatedDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation kontrolü
	if err := updatedDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Gerekli alanlar eksik veya hatalı",
			"details": err.Error(),
		})
		return
	}

	updatedUser := updatedDTO.ToJob()
	var existingUser User
	if err := c.DB.First(&existingUser, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	updatedUser.ID = uint(id)
	result := c.DB.Model(&existingUser).Updates(updatedUser)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User başarıyla güncellendi",
		"data":    updatedUser,
	})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := c.DB.Delete(&User{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
