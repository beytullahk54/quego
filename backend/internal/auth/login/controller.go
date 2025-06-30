package login

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"quego.com/gin-crud/internal/auth/tokens"
	"quego.com/gin-crud/internal/models"
)

type Customer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Controller struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{DB: db}
}

func (c *Controller) Store(r *gin.Context) {
	var data gin.H //Data requestten gelenleri almak için oluşturduğumuz struct

	err2 := r.ShouldBindJSON(&data) //Requestten gelen datayı data structına atıyoruz
	if err2 != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	var property models.User
	result := c.DB.Where("email = ?", data["Email"]).Where("password = ?", data["Password"]).First(&property) //Veritabanından email ile kullanıcı arıyoruz
	if result.Error != nil {
		r.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if property.Password != data["Password"] { //Gelen password ile veritabanındaki password kontrol ediyoruz
		r.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	token, err := tokens.CreateToken(property.Email)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r.JSON(http.StatusOK, gin.H{"token": token, "data": data, "property": property, "message": "Login successful"}) //Başarılı ise response olarak data, property ve message döndürüyoruz
}
