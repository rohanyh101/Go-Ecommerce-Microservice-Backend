package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/roh4nyh/ecom/cmd/api"
	"github.com/roh4nyh/ecom/config"
	"github.com/roh4nyh/ecom/database"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewMySQLStorage(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}

	database.InitDatabase(db)

	server := api.NewApiServer(fmt.Sprintf(":%s", config.Env.Port), db)

	fmt.Printf("Server is running on port :%s\n", config.Env.Port)
	if err := server.Run(); err != nil {
		log.Fatalf("error while running http server: %v", err)
	}
}
