package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func DbConfig() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", db_user, db_password, db_host, db_port, db_name)
	return dsn
}
