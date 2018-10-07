package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "n54l"
	port     = 5432
	user     = "postgres"
	password = "IIvwbWZs7kizRdYwIRot"
	dbname   = "lenslocked_dev"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	db.Close()

	_, err = db.Exec(`
	  INSERT INTO users(name, email)
	  VALUES($1, $2)`,
		"Jon Calhoun", "jon@calhoun.io")
	if err != nil {
		panic(err)
	}

	var id int
	row := db.QueryRow(`
		INSERT INTO users(name, email)
		VALUES($1, $2) RETURNING id`,
		"Jon2 Calhoun2", "jon2@calhoun2.io")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}

}
