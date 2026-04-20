package packageDB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Update() {

	// 🔹 DSN (connection string)
	dsn := "root:Root@1234@tcp(localhost:3306)/db"

	// 🔹 Open DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}
	defer db.Close()

	// 🔹 Check connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	fmt.Println("Connected to database!")

	// 🔹 UPDATE Query
	query := "UPDATE student SET name = ? WHERE roll_no = ?"

	// 🔹 Execute query
	result, err := db.Exec(query, "Tejas", 1)
	if err != nil {
		log.Fatal("Update failed:", err)
	}

	// 🔹 Rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows updated:", rowsAffected)
}
