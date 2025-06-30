package deneme

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func TestFonk(r *gin.Context) {
	customer := []Customer{
		{Name: "Beytullah", Surname: "Kahraman"},
	}

	json, err := json.Marshal(customer)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//deneme := map[string]interface{}{"message": "Hello, World!"}
	r.JSON(http.StatusOK, string(json))
}
