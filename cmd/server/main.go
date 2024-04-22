package main

import (
	"barber/config"
	"barber/database"
	"barber/router"
	"fmt"
)

func main() {
	config.Load()
	database.DBConnect()
	fmt.Println("Running API 🚀")
	router.Generate()
}
