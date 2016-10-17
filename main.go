package main

import (
	"fmt"
	"net/http"
	"database/sql"
)

var db *sql.DB

func main() {
	err := initDB(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected MySql")
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
