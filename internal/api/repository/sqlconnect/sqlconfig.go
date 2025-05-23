package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	fmt.Println("trying to connect to db")
	// connectionString := "root:water@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		// panic(err)
		return nil, err
	}

	fmt.Println("connected to db")

	return db, nil

}
