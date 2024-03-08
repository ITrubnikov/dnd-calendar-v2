package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
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
	r.GET("/albums", FindAlbums)

	r.Run()

	//router := gin.Default()
	//router.GET("/Albums", getAlbums)
	//
	//router.Run("localhost:8080")

}

// GET /books
// Get all books
func FindAlbums(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var albums []model.Album
	db.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"data": albums})
}
