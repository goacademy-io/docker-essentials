package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open(
		"postgres",
		"host="+dbHost+"user="+dbUsername+" password="+dbPassword+" dbname="+dbName+" sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	app := App{
		db: db,
		r:  mux.NewRouter(),
	}
	app.start()
}
