package main

import (
	"fmt"

	"lenslocked.com/models"

	_ "github.com/lib/pq"
)

const (
	host     = "n54l"
	port     = 5432
	user     = "postgres"
	password = "IIvwbWZs7kizRdYwIRot"
	dbname   = "lenslocked_exp"
)

func main() {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	// This will reset the database on evbery run, but is fine for testing
	us.DestructiveReset()

	// Create a user
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dunderfoo.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// Update a user
	user.Name = "Updated Name"
	if err := us.Update(&user); err != nil {
		panic(err)
	}

	// foundUser, err := us.ByEmail("andre.friesen@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(user)

	// // Because of an update, the name should now be "Updated Name"
	// fmt.Prinlln(foundUser)

	// // Delete a user
	// If err:= us.Delete(foundUser.ID); err != nil {
	// 	panic(err)
	// }

	// //Verify the user is deleted
	// _, err= us.ByID(foundUser.ID)
	// if err!= models.ErrNotFound{
	// 	panic("user was deleted")
	// }
}
