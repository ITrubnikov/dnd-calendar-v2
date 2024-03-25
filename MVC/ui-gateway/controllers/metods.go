package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"ui-gateway/model"
)

func getCharacter(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Character{})
}

// GET /books
// Get all books
func FindCharacter(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var characters []model.Character
	db.Find(&characters)
	c.JSON(http.StatusOK, gin.H{"data": characters})
}

func DeleteCharacter(c *gin.Context) {

	log.Println("попытка запуска DeleteAlbum!")
	// Получение объекта базы данных из контекста запроса
	db := c.MustGet("db").(*gorm.DB)
	// Получение ID альбома из параметров запроса
	id := c.Param("id")
	// Выполнение операции удаления
	// Создаем временный объект для хранения информации об удаляемом альбоме (для проверки существования)
	var character model.Character
	if err := db.Where("id = ?", id).First(&character).Error; err != nil {
		// Если альбом с таким ID не найден, отправляем клиенту ошибку 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	// Если альбом найден, производим его удаление
	db.Delete(&character)
	// Отправляем подтверждение об успешном удалении
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// AddAlbum добавляет новый альбом в базу данных.
func AddCharacter(c *gin.Context) {
	// Получение объекта базы данных из контекста запроса
	db := c.MustGet("db").(*gorm.DB)

	// Создание объекта для нового альбома
	var newCharacter model.Character

	// Пробуем получить данные альбома из тела запроса
	if err := c.ShouldBindJSON(&newCharacter); err != nil {
		// В случае ошибки возвращаем сообщение об ошибке клиенту
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавление нового альбома в базу данных
	if err := db.Create(&newCharacter).Error; err != nil {
		// В случае ошибки при добавлении записи возвращаем ошибку сервера
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем подтверждение об успешном добавлении
	c.JSON(http.StatusCreated, gin.H{"data": newCharacter})
}

// UpdateAlbum обновляет существующий альбом в базе данных.
func UpdateCharacter(c *gin.Context) {
	// Получение объекта базы данных из контекста запроса
	db := c.MustGet("db").(*gorm.DB)

	// Получение ID альбома из параметров URL
	id := c.Param("id")

	// Проверяем, существует ли альбом с таким ID
	var character model.Character
	if err := db.Where("id = ?", id).First(&character).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	// Получаем обновленные данные альбома из тела запроса
	var updatedData model.Character
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновляем альбом в базе данных
	if err := db.Model(&character).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправляем подтверждение об успешном обновлении
	c.JSON(http.StatusOK, gin.H{"data": character})
}
