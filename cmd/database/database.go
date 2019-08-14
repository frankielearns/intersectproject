package lib

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Databaseconnect(stocktick string, price float64, username string) {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO stockhistory (stocktick, price, username)
VALUES ($1, $2, $3)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, stocktick, price, username).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
