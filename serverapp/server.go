package main

import (
	"serverapp/config"
	"database/sql"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	_ "github.com/lib/pq"
)

var (
	// Changed variable name from 'config' to 'configVar'
	configVar = config.LoadConfig()
	db        *sql.DB
)

func main() {
	flag.Parse()

	// Connect to PostgreSQL.
	connStr := fmt.Sprintf("%s/%s", configVar.DBURI, configVar.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Register handlers.
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/version", versionHandler)

	log.Printf("Server is listening on http://%s\n", configVar.Addr)
	log.Fatal(http.ListenAndServe(configVar.Addr, nil))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Guest"
	}

	// Insert the greeting into PostgreSQL.
	_, err := db.Exec("INSERT INTO "+configVar.Collection+" (name, greeting) VALUES ($1, $2)", name, configVar.Greeting)
	if err != nil {
		log.Printf("Error inserting greeting into PostgreSQL: %v", err)
	}

	// Retrieve greetings from PostgreSQL.
	var greeting string
	err = db.QueryRow("SELECT greeting FROM "+configVar.Collection+" WHERE name = $1", name).Scan(&greeting)
	if err != nil {
		log.Printf("Error querying PostgreSQL: %v", err)
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "%s, %s!\n", greeting, html.EscapeString(name))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

