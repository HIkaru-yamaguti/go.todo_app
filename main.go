package main

import (
	"fmt"
	"go.todo_app/app/controllers"
	"go.todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}