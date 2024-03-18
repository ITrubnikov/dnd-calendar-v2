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
	r.POST("/albums", AddAlbum)
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

// AddAlbum добавляет новый альбом в базу данных.
func AddAlbum(c *gin.Context) {
	// Получение объекта базы данных из контекста запроса
	db := c.MustGet("db").(*gorm.DB)

	// Создание объекта для нового альбома
	var newAlbum model.Album

	// Пробуем получить данные альбома из тела запроса
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		// В случае ошибки возвращаем сообщение об ошибке клиенту
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавление нового альбома в базу данных
	if err := db.Create(&newAlbum).Error; err != nil {
		// В случае ошибки при добавлении записи возвращаем ошибку сервера
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем подтверждение об успешном добавлении
	c.JSON(http.StatusCreated, gin.H{"data": newAlbum})
}
