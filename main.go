package main

import (
	"fmt"
	"log"
	"net/http"
	"quizofkingz/entity"
	"quizofkingz/handlers"
	"quizofkingz/repository/mysql"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	http.HandleFunc("/users/register", handlers.UserRegisterHandler)
	http.HandleFunc("/users/login", handlers.UserLoginHandler)
	http.HandleFunc("/users/profile", handlers.GetProfile)

	http.ListenAndServe(":8080", nil)

	fmt.Printf("Server started at port 8080\n")
}

func testUserMysqlRepo() {
	mysqlRepo, err := mysql.New()
	if err != nil {
		log.Fatal(err)
	}

	createdUser, err := mysqlRepo.RegisterUser(entity.User{PhoneNumber: "09197321182", Name: "Alireza", Password: "1234"})

	if err != nil {
		fmt.Println("register user", err)
	} else {
		fmt.Println("created user", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)

	if err != nil {
		fmt.Println("unique err", err)
	}

	fmt.Println("isUnique", isUnique)
}
