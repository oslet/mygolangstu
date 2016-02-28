package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func insert(db *sql.DB) {
	stmt, err := db.Prepare("insert into url(status, timestamp) values(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("200", 1456665246)
	stmt.Exec("500", 1456665720)
}
func main() {
	db, err := sql.Open("mysql", "root:eletmc@tcp(127.0.0.1:3306)/checkurl?charset=utf8")
	if err != nil {
		log.Fatal("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	insert(db)
	rows, err := db.Query("select * from url")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	fmt.Println("")
	cols, _ := rows.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println("")
}
