package packageDB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Insert() {

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

	// 🔹 INSERT Query
	query := "INSERT INTO student (roll_no, name) VALUES (?, ?)"

	// 🔹 Execute query
	result, err := db.Exec(query, 10, "John Doe")
	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	// 🔹 Rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows Inserted:", rowsAffected)
}
