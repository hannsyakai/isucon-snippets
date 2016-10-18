package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	err := initDB(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected MySql")

	http.HandleFunc("/hello", helloHandler)
	initStaticFiles("assets", "assets")

	nginxSockPath := os.Getenv("NGINX_UNIX_DOMAIN_SOCK_PATH")

	if nginxSockPath == "" {
		http.ListenAndServe(":8080", nil)
	} else {
		os.Remove(nginxSockPath)
		ul, err := net.Listen("unix", nginxSockPath)
		if err != nil {
			panic(err)
		}
		os.Chmod(nginxSockPath, 0777)
		defer ul.Close()
		log.Fatal(http.Serve(ul, nil))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
