package main

import (
	"database/sql"
	"files2pg/internal/config"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "testpwd"
	dbname   = "test"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		panic(err)
	} else {
		fmt.Println(config)
		fmt.Println(config.Db.Password)
		fmt.Println(config.Db.Address)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}
