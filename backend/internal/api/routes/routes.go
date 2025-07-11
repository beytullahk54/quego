package routes

import (
	login "quego.com/gin-crud/internal/auth/login"
	deneme "quego.com/gin-crud/internal/deneme"
	jobs "quego.com/gin-crud/internal/jobs"
	users "quego.com/gin-crud/internal/users"

	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var logFile *os.File

func init() {
	var err error
	logFile, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
}

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Cron job'ı başlat
	c := cron.New()
	jobController := jobs.NewController(db)

	// Her 1 dakikada bir çalışacak cron job

	_, err := c.AddFunc("*/1 * * * *", func() {
		//log.Println("Jobs kontrol ediliyor...")
		jobController.CheckAndProcessJobs()
	})

	if err != nil {
		log.Printf("Cron job başlatılırken hata oluştu: %v", err)
	}
	c.Start()

	// User routes
	UserController := users.NewController(db)
	r.GET("/users", UserController.GetUsers)
	r.GET("/users/:id", UserController.GetUserByID)
	r.POST("/users", UserController.CreateUser)
	r.PUT("/users/:id", UserController.UpdateUser)
	r.DELETE("/users/:id", UserController.DeleteUser)

	// Property routes
	r.GET("/jobs", jobController.GetJobs)
	r.GET("/jobs/:id", jobController.GetJobByID)
	r.POST("/jobs", jobController.CreateJob)
	r.PUT("/jobs/:id", jobController.UpdateJob)
	r.DELETE("/jobs/:id", jobController.DeleteJob)

	r.GET("/deneme", deneme.TestFonk)

	loginController := login.NewController(db)
	r.POST("/login", loginController.Store)

	r.GET("/test", func(c *gin.Context) {
		log.Println("Test endpoint çağrıldı")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
}
