package main

import (
	"fmt"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	host     string
	user     string
	pass     string
	port     int
	database string
	db       *sql.DB
)

func init() {
	host = "127.0.0.1"
	user = "root"
	pass = "mysql"
	port = 3306
	database = "mysql"
}

func dbSource(user, pass, address, database string, port int, protocol string) string {
	if address == "localhost" {
		address = ""
	}

	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, pass, protocol, address, port, database)
}

func sample(w http.ResponseWriter, r *http.Request) {
	fmt.Println(dbSource(user, pass, host, database, port, "tcp"))

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()

	var err error
	db, err = sql.Open("mysql", dbSource(user, pass, host, database, port, "tcp"))
	if err != nil {
		panic(errors.Wrap(err, "sql.Open error."))
	}
	defer db.Close()

	var v []uint8
	sql := `
SELECT version();
`
	err = db.QueryRow(sql).Scan(&v)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(v))
}

func main() {
	http.HandleFunc("/", sample)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
