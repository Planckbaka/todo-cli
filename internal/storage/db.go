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

func GetAllTodoData() ([]models.Todo, error) {
	var todos []models.Todo
	result := dbConn.Model(&models.Todo{}).Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func DeleteTodoData(id string) (models.Todo, error) {
	var todo models.Todo
	dbConn.Model(&models.Todo{}).Where("id = ?", id).First(&todo)
	result := dbConn.Where("id = ?", id).Delete(&models.Todo{})
	if result.Error != nil {
		return models.Todo{}, result.Error
	}
	return todo, nil
}

// DoneTodoData signature a task done
func DoneTodoData(id string) error {
	result := dbConn.Model(&models.Todo{}).Where("id = ?", id).Update("completed", true)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// QueryTodoData find data by keywords
func QueryTodoData(title string) ([]models.Todo, error) {
	var todos []models.Todo
	result := dbConn.Model(&models.Todo{}).
		Where("title LIKE ?", "%"+title+"%").
		Find(&todos).
		Limit(5)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
