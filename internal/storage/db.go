package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/Planckbaka/todo-cli/internal/config"
	"github.com/Planckbaka/todo-cli/internal/errors"
	"github.com/Planckbaka/todo-cli/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConn *gorm.DB

func InitDatabase() error {
	configs := config.Load()

	if err := os.MkdirAll(configs.DatabasePathDir, 0755); err != nil {
		return fmt.Errorf("could not create database directory: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(configs.DatabasePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDatabaseConnection, err)
	}

	// set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	dbConn = db

	// Migrate the schema
	return db.AutoMigrate(&models.Todo{})
}

// close database function
//func CloseDatabaseElegantly(ctx context.Context) error {
//	if dbConn == nil {
//		return nil
//	}
//
//	sqlDB, err := dbConn.DB()
//	if err != nil {
//		return err
//	}
//
//	done := make(chan error, 1)
//	go func() {
//		done <- sqlDB.Close()
//	}()
//
//	select {
//	case <-ctx.Done():
//		// 如果超时或被取消
//		return ctx.Err()
//	case err := <-done:
//		return err
//	}
//}

func CloseDatabase() error {
	if dbConn != nil {
		sqlDB, err := dbConn.DB()
		if err != nil {
			return err
		}
		log.Print("Close database successfully")
		return sqlDB.Close()
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
	result := dbConn.Model(&models.Todo{}).Order("priority asc").Find(&todos)
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
	configs := config.Load()

	var todos []models.Todo
	result := dbConn.Model(&models.Todo{}).
		Where("title LIKE ?", "%"+title+"%").
		Find(&todos).
		Limit(configs.MaxQueryResults)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
