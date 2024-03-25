package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	metod "ui-gateway/controllers"
	"ui-gateway/model"
)

func main() {
	r := gin.Default()
	db := model.SetupModels() // new

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	//// GET /books
	//// Get all books
	r.Use(cors.Default())
	r.GET("/Character", metod.FindCharacter)
	r.DELETE("/Character/:id", metod.DeleteCharacter)
	r.POST("/Character", metod.AddCharacter)
	r.PUT("/Character/:id", metod.UpdateCharacter) // ИЛИ r.PATCH("/albums/:id", handlers.UpdateAlbum)
	r.Run()

	//router := gin.Default()
	//router.GET("/Albums", getAlbums)
	//
	//router.Run("localhost:8080")

}
