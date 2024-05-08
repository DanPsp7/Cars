package repository

import (
	"database/sql"
	"log"

	"github.com/Cars/internal/models"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func newDB(connectionString string) (*DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &DB{conn: db}, nil
}

// SaveCar - сохранение данных автомобиля в БД
func (db *DB) SaveCar(car models.Car) error {
	_, err := db.conn.Exec("INSERT INTO cars (reg_num, mark, model, year) VALUES ($1, $2, $3, $4)",
		car.RegNum, car.Mark, car.Model, car.Year)
	if err != nil {
		log.Println("Ошибка при сохранении автомобиля:", err)
		return err
	}
	return nil
}

// GetCarByID - получение данных автомобиля по идентификатору
func (db *DB) GetCarByID(id int) (models.Car, error) {
	var car models.Car
	err := db.conn.QueryRow("SELECT reg_num, mark, model, year FROM cars WHERE id = $1", id).
		Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year)
	if err != nil {
		log.Println("Ошибка при получении данных автомобиля:", err)
		return models.Car{}, err
	}
	return car, nil
}

func (db *DB) GetWithFilter(regNum, mark, model, owner string) ([]models.Car, error) {
	query := "SELECT reg_num, mark, model,year FRROM cars WHERE 1=1"
	args := []interface{}{}
	if regNum != "" {
		query += " AND reg_num = $1"
		args = append(args, regNum)
	}
	if mark != "" {
		query += " AND mark = $2"
		args = append(args, mark)
	}
	if model != "" {
		query += " AND model = $3"
		args = append(args, model)
	}
	if owner != "" {
		query += " AND owner = $4"
		args = append(args, owner)
	}
	rows, err := db.conn.Query(query, args...)
	if err != nil {
		log.Println("ошибка при выполнении запроса", err)
		return nil, err
	}
	defer rows.Close()
	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year)
		if err != nil {
			log.Println("Ошибка при сканировании данных:", err)
			continue
		}
		cars = append(cars, car)
	}

	return cars, nil

}

func (db *DB) GetAllCar() ([]models.Car, error) {
	rows, err := db.conn.Query("SELECT reg_num, mark, modek , year FROM cars")
	if err != nil {
		log.Fatal(nil, err)

	}
	defer rows.Close()
	var cars []models.Car

	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year)
		if err != nil {
			log.Fatal(err)
			continue
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (db *DB) UpdateCarByID(id int, updatedCar models.Car) error {
	_, err := db.conn.Exec("UPDATE cars SET reg_num = $1, mark = $2, model = $3, year = $4 WHERE id = $5",
		updatedCar.RegNum, updatedCar.Mark, updatedCar.Model, updatedCar.Year, id)
	if err != nil {
		log.Println("Ошибка при обновлении данных автомобиля:", err)
		return err
	}
	return nil
}

func (db *DB) DeleteCarByID(id int) error {
	_, err := db.conn.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка при удалении автомобиля:", err)
		return err
	}
	return nil
}

// Закрытие соединения с БД
func (db *DB) Close() {
	db.conn.Close()
}
