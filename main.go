package main

import (
	"bwacoba/users/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwacoba?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userinput := user.RegisterUserInput{}
	userinput.Name = "Fredio"
	userinput.Occupation = "DevOps"
	userinput.Email = "onamel313@gmail.com"
	userinput.Password = "password"
	userService.RegisterUser(userinput)
}
