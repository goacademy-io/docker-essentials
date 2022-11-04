package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"os"
)

func main() {
	pass := os.Getenv("DB_PASS")
	db, err := gorm.Open(
		"postgres",
		"host=students-db user=postgres password="+pass+" dbname=students sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	app := App{
		db: db,
		r:  mux.NewRouter(),
	}
	app.start()
}
