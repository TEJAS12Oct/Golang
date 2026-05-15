// repository/user_repo.go
package repository

import (
	"RESTAPIPOSTGIN/model"
	"database/sql"
	"fmt"
)

type UserRepo interface {
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
}

type userRepoImpl struct { // userRepoImpl struct has a field db of type *sql.DB,
	// which is a pointer to a sql.DB struct that represents a connection to the database
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo { // NewUserRepo is a constructor function that takes a pointer to a sql.DB as an argument
	// and returns a UserRepo
	return &userRepoImpl{db: db} // it initializes the db field of the userRepoImpl struct
	// with the provided sql.DB and returns a pointer to the userRepoImpl struct
}

func (r *userRepoImpl) GetUsers() ([]model.User, error) { // GetUsers is a method of the userRepoImpl struct that
	// implements the GetUsers method of the UserRepo interface
	rows, err := r.db.Query("SELECT id, name FROM users") // it executes a SQL query to select the id and name columns from the users table
	if err != nil {                                       //	 if there is an error executing the query, it returns nil and the error
		return nil, err // it returns nil for the slice of users and the error that occurred
	}
	defer rows.Close() // it defers the closing of the rows until the function returns,
	// to ensure that the resources are properly released

	var users []model.User // it initializes an empty slice of model.User to store the users retrieved from the database
	for rows.Next() {      // it iterates over the rows returned by the query using the Next method of the rows object
		var u model.User          // it declares a variable u of type model.User to store the data for each user retrieved from the database
		rows.Scan(&u.ID, &u.Name) // it scans the values of the id and name columns from the current row into the ID and Name fields of the user variable u
		users = append(users, u)  // it appends the user variable u to the users slice, which will contain all the users retrieved from the database
	}
	return users, nil // it returns the slice of users and nil for the error, indicating that the operation was successful
}

func (r *userRepoImpl) CreateUser(user model.User) (model.User, error) {
	var id int

	err := r.db.QueryRow(
		"INSERT INTO users (name) VALUES ($1) RETURNING id",
		user.Name,
	).Scan(&id)

	if err != nil {
		return model.User{}, err
	}

	user.ID = id
	return user, nil
}

func (r *userRepoImpl) UpdateUser(user model.User) (model.User, error) {
	var id int

	err := r.db.QueryRow(
		"UPDATE users SET name = $1 WHERE id = $2 RETURNING id",
		user.Name, user.ID,
	).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}

	user.ID = id
	fmt.Println("Repo received:", user)
	return user, nil
}

func (r *userRepoImpl) DeleteUser(id int) error {

	result, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil

}
