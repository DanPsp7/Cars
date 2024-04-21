package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Структура для хранения данных об автомобиле
type Car struct {
	ID     uint   `gorm:"primaryKey"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
}

// Структура для хранения данных о владельце автомобиля
type People struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&Car{}, &People{})

	r := gin.Default()

	r.GET("/cars", getCars)
	r.DELETE("/cars/:id", deleteCar)
	r.PUT("/cars/:id", updateCar)
	r.POST("/cars", addCars)

	r.Run(":8080")
}

func getCars(c *gin.Context) {
	// Реализация фильтрации и пагинации
	// Возвращаем список автомобилей
}

func deleteCar(c *gin.Context) {
	id := c.Param("id")
	// Реализация удаления автомобиля по id
}

func updateCar(c *gin.Context) {
	id := c.Param("id")
	// Реализация изменения полей автомобиля по id
}

func addCars(c *gin.Context) {
	var data struct {
		RegNums []string `json:"regNums"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, regNum := range data.RegNums {
		response, err := http.Get(fmt.Sprintf("http://external-api/info?regNum=%s", regNum))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer response.Body.Close()

		var carData Car
		if err := json.NewDecoder(response.Body).Decode(&carData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Сохранение данных в БД
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cars added"})
}
