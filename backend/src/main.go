package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"

	"trackr/src/controllers"
	"trackr/src/db"
)

func main() {
	godotenv.Load()

	serviceProvider := db.InitServiceProvider(sqlite.Open("bin/test.db"))
	router := controllers.InitRouter(serviceProvider)
	router.Run()
}
