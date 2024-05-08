package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация сервера
	port := 8080
	fmt.Printf("Сервер запущен на порту %d\n", port)

	// Настройка маршрутов

	router := gin.Default()
	router.GET("/api", getCars)
	router.POST("/api", updateCar)
	router.DELETE()
	router.PUT()

	// Запуск приложения
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

	db, err := repository.NewDB("user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal("Ошибка при подключении к БД:", err)
	}
	defer db.Close()

	// Выполнение миграций
	_, err = db.conn.Exec(migrationSQL)
	if err != nil {
		log.Fatal("Ошибка при выполнении миграций:", err)
	}

}
