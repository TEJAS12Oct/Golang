package Repository

import (
	db "JWT/db"
	models "JWT/models"
)

func CreateUser(user models.User) error {
	_, err := db.DB.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username, user.Password,
	)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	err := db.DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username=?",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
