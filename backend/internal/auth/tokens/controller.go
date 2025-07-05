package tokens

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"quego.com/gin-crud/config"
	"quego.com/gin-crud/internal/models"
)

var secretKey = []byte("secret-key")
var db *gorm.DB

func init() {
	db = config.InitDB()
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	var user models.User
	db.First(&user, "email = ?", username)

	result := db.Create(&models.Token{UserID: strconv.Itoa(int(user.ID)), TokenID: tokenString})
	if result.Error != nil {
		return "", result.Error
	}

	return tokenString, nil
}

func VerifyToken(ctx *gin.Context) error {
	token := ctx.GetHeader("Authorization")[7:]
	/*token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})*/

	var tokens models.Token
	db.First(&tokens, "token_id = ?", token)

	if tokens.TokenID == "" {
		return fmt.Errorf("token not found")
	}

	return nil
}
