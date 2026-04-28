package Repository

import (
	db "RestAPIDatabaseConnection/DB"
	models "RestAPIDatabaseConnection/Models"
	"fmt"
)

func GetUsers() ([]models.Student, error) {
	rows, err := db.DB.Query("SELECT id, name FROM student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var s models.Student
		err := rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

func InsertUser(name string) error {
	query := "INSERT INTO student (name) VALUES (?)"
	_, err := db.DB.Exec(query, name)
	return err
}

func UpdateUser(id int, name string) error {
	query := "UPDATE student SET name=? WHERE id=?"

	result, err := db.DB.Exec(query, name, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no record found with id %d", id)
	}

	return nil
}

func DeleteUser(id int) error {
	query := "DELETE FROM student WHERE id=?"
	_, err := db.DB.Exec(query, id)
	return err
}
