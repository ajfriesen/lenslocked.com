package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 1; i <= 6; i++ {
		userID := 1
		if i > 3 {
			userID = 2
		}
		amount := i * 100
		description := fmt.Sprintf("USB-C Adapter x%d", i)

		rows, err := db.Query(`
			SELECT * FROM users
			INNER JOIN orders ON users.id=orders.user_id
			INSERT INTO orders(user_id, amount, description)
			VALUES($1,$2,$3)`, userID, amount, description)
		if err != nil {
			panic(err)
		}
	}
}
