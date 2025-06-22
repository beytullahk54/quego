package main

import (
	"fmt"
	"log"

	"quego.com/gin-crud/config"
	"quego.com/gin-crud/internal/api/routes"
	"quego.com/gin-crud/internal/auth/tokens"
	"quego.com/gin-crud/internal/jobs"
	"quego.com/gin-crud/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Sunucu başlatılıyor...")
	fmt.Println("http://localhost:8080 adresinde çalışıyor")

	// Veritabanı bağlantısı
	db := config.InitDB()

	// Tabloları oluştur
	if err := db.AutoMigrate(&users.User{}); err != nil {
		log.Fatal("Tablolar oluşturulamadı:", err)
	}
	if err := db.AutoMigrate(&jobs.Job{}); err != nil {
		log.Fatal("Tablolar oluşturulamadı:", err)
	}
	if err := db.AutoMigrate(&tokens.Token{}); err != nil {
		log.Fatal("Tablolar oluşturulamadı:", err)
	}

	// Gin ayarları
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Route'ları ayarla
	routes.SetupRoutes(r, db)

	// Sunucuyu başlat
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Sunucu başlatılamadı:", err)
	}
}
