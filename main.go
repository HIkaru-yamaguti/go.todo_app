package main

import (
	"fmt"
	//"os/user"

	//"go.todo_app/app/controllers"
	"go.todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	//controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)
}