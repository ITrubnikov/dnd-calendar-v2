package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
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
	r.DELETE("/albums/:id", DeleteAlbum)
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

func DeleteAlbum(c *gin.Context) {

	log.Println("попытка запуска DeleteAlbum!")
	// Получение объекта базы данных из контекста запроса
	db := c.MustGet("db").(*gorm.DB)
	// Получение ID альбома из параметров запроса
	id := c.Param("id")
	// Выполнение операции удаления
	// Создаем временный объект для хранения информации об удаляемом альбоме (для проверки существования)
	var album model.Album
	if err := db.Where("id = ?", id).First(&album).Error; err != nil {
		// Если альбом с таким ID не найден, отправляем клиенту ошибку 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	// Если альбом найден, производим его удаление
	db.Delete(&album)
	// Отправляем подтверждение об успешном удалении
	c.JSON(http.StatusOK, gin.H{"data": true})
}
