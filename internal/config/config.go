package config

import (
	"errors"
	"log"

	"github.com/Planckbaka/todo-cli/internal/models"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() error {
	// setting config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	// default value
	viper.SetDefault("database.path", "./data/todos.db")
	viper.SetDefault("database.dirPath", "./data")
	viper.SetDefault("max.queryResults", 5)
	// read config
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
		log.Println("未找到配置文件,使用默认值和环境变量,自动写入配置文件地址'./configs/config.yaml'")
		err := viper.SafeWriteConfigAs("./configs/config.yaml")
		if err != nil {
			//log.Printf("configs file already exists:%v", err)
		}
	}

	// 解析到结构体
	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	// 配置热更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件变更: %s", e.Name)
		err := viper.Unmarshal(&config)
		if err != nil {
			return
		}
	})
	return nil
}

func Load() *models.Config {
	// 支持环境变量和配置文件
	return &models.Config{
		DatabasePath:    viper.GetString("database.path"),
		DatabasePathDir: viper.GetString("database.dirPath"),
		MaxQueryResults: 5,
	}
}
