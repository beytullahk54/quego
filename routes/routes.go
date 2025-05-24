package routes

import (
	jobs "quego.com/gin-crud/jobs"
	users "quego.com/gin-crud/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// User routes
	UserController := users.NewController(db)
	r.GET("/users", UserController.GetUsers)
	r.GET("/users/:id", UserController.GetUserByID)
	r.POST("/users", UserController.CreateUser)
	r.PUT("/users/:id", UserController.UpdateUser)
	r.DELETE("/users/:id", UserController.DeleteUser)

	// Property routes
	jobController := jobs.NewController(db)
	r.GET("/jobs", jobController.GetJobs)
	r.GET("/jobs/:id", jobController.GetJobByID)
	r.POST("/jobs", jobController.CreateJob)
	r.PUT("/jobs/:id", jobController.UpdateJob)
	r.DELETE("/jobs/:id", jobController.DeleteJob)
}
