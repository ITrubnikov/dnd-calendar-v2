package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	// Enable VIPER to read Environment Variables

	// To get the value from the config file using key
	// viper package read .env
	viper_user := "mygouser"
	viper_password := "mygouser"
	viper_db := "postgres"
	viper_host := "postgres"
	viper_port := "5432"
	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Album{})
	// Initialise value
	//m := Album{ID: "6", Title: "author1", Artist: "title1", Price: 23}
	//db.Create(&m)
	return db
}
