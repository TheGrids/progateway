package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"progateway.ru/m/v2/models"
	"progateway.ru/m/v2/services"
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

	api := r.Group("/api")
	{
		{
			api.GET("/alltest", services.GetAllTest)
			api.GET("/testabc", services.AddTest) //временно
			api.GET("/test/:id", services.Test)
			api.POST("/submit/:testID", services.SubmitTest)
		}
		auth := api.Group("/auth")
		{
			auth.POST("/signup", services.SignUp)
			auth.POST("/signin", services.SignIn)
			auth.GET("/logout", services.Logout)
			api.GET("/activate/:uuid", services.Activate)
		}
	}

	err := r.Run()
	if err != nil {
		return
	}
}
