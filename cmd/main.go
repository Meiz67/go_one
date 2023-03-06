package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Response struct {
	Version string `json:"Version"`
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "port=5432 host=postgres user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	query := "select now(), version();"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	var (
		time, version string
	)
	for rows.Next() {
		err := rows.Scan(&time, &version)
		if err != nil {
			continue
		}
	}
	response, err := json.Marshal(Response{Version: version})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/version", versionHandler)
	http.Handle("/", router)
	fmt.Println("Listening :8040")
	err := http.ListenAndServe(":8040", nil)
	if err != nil {
		log.Fatal(err)
	}
}
