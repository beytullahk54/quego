package tokens

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"quego.com/gin-crud/config"
	"quego.com/gin-crud/internal/users"
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

	var user users.User
	db.First(&user, "email = ?", username)

	result := db.Create(&Token{UserID: strconv.Itoa(int(user.ID)), TokenID: tokenString})
	if result.Error != nil {
		return "", result.Error
	}

	return tokenString, nil
}

/*func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}*/
