package config

import (
	"Go_Postrgres/API/structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	//Set database, harus dibuat dahulu nama databasenya
	//Format args yang digunakan: username:password@tcp(127.0.0.1:3306)/namadatabase?charset=utf8&parseTime=True
	/*db, err := gorm.Open("pq", "postgres:root@tcp(127.0.0.1:3306)/go_postgres?charset=utf8&parseTime=True")*/
	dsn := "host=localhost user=postgres password=root dbname=go_postgres port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
