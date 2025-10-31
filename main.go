/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"

	"github.com/Planckbaka/todo-cli/cmd"
	"github.com/Planckbaka/todo-cli/internal/storage"
)

func main() {
	err := storage.InitDatabase()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Execute()

	//close database
	defer func() {
		err := storage.CloseDatabase()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//// close elegantly
	//// 优雅关闭
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//
	//log.Println("\n🛑 收到关闭信号,优雅关闭中...")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := storage.CloseDatabaseElegantly(ctx); err != nil {
	//	log.Printf("数据库关闭失败: %v\n", err)
	//
	//}
	//log.Println("👋 服务已关闭")

}
