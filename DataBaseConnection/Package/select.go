package packageDB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Select() {

	dsn := "root:Root@1234@tcp(localhost:3306)/db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Connection opened successfully!", err)
		log.Fatal("Error opening connection:", err)
		panic(err)
	}
	defer db.Close()

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	fmt.Println("Connected to database!")

	// Query
	rows, err := db.Query("SELECT roll_no, name FROM student")
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatal("Query error:", err)
	}
	defer rows.Close()

	// Iterate rows
	for rows.Next() {
		var rollNo int
		var name string

		err := rows.Scan(&rollNo, &name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rollNo, name)
	}

	// Check for errors after loop
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
