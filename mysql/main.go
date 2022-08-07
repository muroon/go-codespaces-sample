package main

import (
	"flag"
	"fmt"

	"database/sql"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	host     string
	user     string
	pass     string
	port     int
	database string
	db       *sql.DB
	dryRun   *bool
)


func init() {
	host = "127.0.0.1"
	user = "root"
	pass = ""
	port = 3306
	database = "mysql"
	dryRun = flag.Bool("dry_run", true, "usesage of dry run")
}

func dbSource(user, pass, address, database string, port int, protocol string) string {
	if address == "localhost" {
		address = ""
	}

	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, pass, protocol, address, port, database)
}

func main() {
	flag.Parse()
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

	fmt.Println(string(v))
}
