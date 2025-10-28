package storage

import (
	"github.com/Planckbaka/todo-cli/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("./data/todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbConn = db
	// Migrate the schema
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}
	return nil
}

func SaveTodoData(todo *models.Todo) error {
	result := dbConn.Model(&models.Todo{}).Create(todo)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func QueryTodoData() ([]models.Todo, error) {
	var todos []models.Todo
	dbConn.Model(&models.Todo{}).Find(&todos)
	return todos, nil
}
