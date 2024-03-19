package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	metod "ui-gateway/controllers"
	"ui-gateway/model"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Album{})
}

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
	r.GET("/albums", metod.FindAlbums)
	r.DELETE("/albums/:id", metod.DeleteAlbum)
	r.POST("/albums", metod.AddAlbum)
	r.PUT("/albums/:id", metod.UpdateAlbum) // ИЛИ r.PATCH("/albums/:id", handlers.UpdateAlbum)
	r.Run()

	//router := gin.Default()
	//router.GET("/Albums", getAlbums)
	//
	//router.Run("localhost:8080")

}
