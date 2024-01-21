package main

import (
	"context"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI  = flag.String("mongoURI", "mongodb://localhost:27017", "MongoDB URI")
	greeting  = flag.String("g", "Hello", "Greet with `greeting`")
	addr      = flag.String("addr", "0.0.0.0:8080", "address to serve")
	collection *mongo.Collection
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: service [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Connect to MongoDB.
	client, err := mongo.NewClient(options.Client().ApplyURI(*mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Select the database and collection.
	database := client.Database("mydatabase")
	collection = database.Collection("greetings")

	// Register handlers.
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/version", versionHandler)

	log.Printf("Server is listening on http://%s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Guest"
	}

	// Insert the greeting into MongoDB.
	_, err := collection.InsertOne(context.Background(), map[string]interface{}{"name": name, "greeting": *greeting})
	if err != nil {
		log.Printf("Error inserting greeting into MongoDB: %v", err)
	}

	// Retrieve greetings from MongoDB.
	var result map[string]interface{}
	err = collection.FindOne(context.Background(), map[string]interface{}{"name": name}).Decode(&result)
	if err != nil {
		log.Printf("Error querying MongoDB: %v", err)
	}

	greeting := result["greeting"].(string)

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

