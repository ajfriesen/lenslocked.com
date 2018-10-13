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

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Opening database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Ping the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// The first version w/out the ID
	// _, err = db.Exec(`
	//   INSERT INTO users(name, email)
	//   VALUES($1, $2)`,
	// 	"Jon Calhoun", "jon@calhoun.io")
	// if err != nil {
	// 	panic(err)
	// }

	// // The second version that returns the ID
	// var id int
	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id`,
	// 	"Andrej Friesen", "andre.friesen@gmail.com")
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// var id int
	// var name, email string
	// rows, err := db.Query(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE id > $1`,
	// 	0)
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	rows.Scan(&id, &name, &email)
	// 	fmt.Println("ID:", id, "Name:", name, "Email:", email)
	// }

	// var id int
	// var name, email string
	// row := db.QueryRow(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE id=$1`, 4)
	// err = row.Scan(&id, &name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("ID:", id, "Name:", name, "Email:", email)

	// var id int
	// var name, email string
	// rows, err := db.Query(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE id > $1`, 0)
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	rows.Scan(&id, &name, &email)
	// 	fmt.Println("ID:", id, "Name:", name, "Email:", email)
	// }

	var id int
	for i := 1; i < 6; i++ {
		//create some fake data
		userId := 1
		if i > 3 {
			userId = 2
		}
		amount := 1000 * i
		description := fmt.Sprintf("USB-C Adapter x%d", i)

		err = db.QueryRow(`
			INSERT INTO orders (user_id, amount, description)
			VALUES ($1,$2,$3)
			RETURNING id`,
			userId, amount, description).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Created an order with ID:", id)
	}

	db.Close()

}
