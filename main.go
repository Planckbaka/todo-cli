/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/Planckbaka/todo-cli/cmd"
	"github.com/Planckbaka/todo-cli/internal/storage"
)

func main() {
	err := storage.InitDatabase()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Execute()
}
