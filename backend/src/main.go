package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"

	"trackr/src/controllers"
	"trackr/src/services_impl"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables from .env file:", err)
		return
	}

	var dialector gorm.Dialector
	if os.Getenv("DB_TYPE") == "sqlite" {
		dialector = sqlite.Open(os.Getenv("DB_CONNECTION_STRING"))
	} else if os.Getenv("DB_TYPE") == "mysql" {
		dialector = mysql.New(mysql.Config{
			DSN:               os.Getenv("DB_CONNECTION_STRING"),
			DefaultStringSize: 256,
		})
	} else {
		panic("Unsupported database type")
	}

	if os.Getenv("MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	serviceProvider := services_impl.InitServiceProvider(dialector)
	router := controllers.InitRouter(serviceProvider)

	if err := router.Run(); err != nil {
		fmt.Println("Error running router:", err)
		return
	}
}
