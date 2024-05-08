package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Cars/internal/models"
	"github.com/Cars/internal/repository"
)

var db *repository.DB

// GET /api/v1/cars
func getCars(w http.ResponseWriter, r *http.Request) {
	// Получение параметров фильтрации из URL-параметров
	regNum := r.URL.Query().Get("regNum")
	mark := r.URL.Query().Get("mark")
	model := r.URL.Query().Get("model")
	owner := r.URL.Query().Get("owner")

	// Вызов метода для получения данных с применением фильтров
	cars, err := db.GetWithFilter(regNum, mark, model, owner)
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	// Преобразование в JSON и отправка клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	// Логика удаления автомобиля
	// Пример: Получение идентификатора автомобиля из URL параметров
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Ошибка при преобразовании идентификатора", http.StatusBadRequest)
		return
	}
	err = db.DeleteCarByID(intID)
	if err != nil {
		http.Error(w, "Ошибка при удалении автомобиля", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Автомобиль успешно удален"))
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedCar models.Car
	err := json.NewDecoder(r.Body).Decode(&updatedCar)
	if err != nil {
		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
		return
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Ошибка при преобразовании идентификатора", http.StatusBadRequest)
		return
	}
	err = db.UpdateCarByID(intID, updatedCar)
	if err != nil {
		http.Error(w, "Ошибка при обновлении автомобиля", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные автомобиля успешно обновлены"))
}

func addCar(w http.ResponseWriter, r *http.Request) {
	// Логика добавления новых автомобилей
	// Пример: Распаковка JSON-данных из запроса
	var newCar models.Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
		return
	}
	externalAPIURL := "https://localhost:8080/cars" // Замените на реальный URL вашего внешнего API
	requestBody, _ := json.Marshal(newCar)          // Преобразование нового автомобиля в JSON
	response, err := http.Post(externalAPIURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Ошибка при обращении к внешнему API", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	// Запрос к внешнему API (описанному в Swagger-документе)
	// ...

	// Сохранение данных в БД
	err = db.SaveCar(newCar)
	if err != nil {
		http.Error(w, "Ошибка при сохранении автомобиля", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Автомобиль успешно добавлен"))
}
