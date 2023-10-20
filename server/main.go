package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"progateway.ru/m/v2/models"
	"time"
)

func main() {
	r := gin.Default()

	models.ConnectionDataBase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://progateway.ru", "http://127.0.0.1:5173"}, //на проде убрать localhost
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           1 * time.Minute,
	}))

	err := r.Run()
	if err != nil {
		return
	}
}
