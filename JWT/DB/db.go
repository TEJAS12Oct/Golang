package DB

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "root:Root@1234@tcp(127.0.0.1:3306)/GoLRestDatabaseConnection"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ DB Connected")
}
