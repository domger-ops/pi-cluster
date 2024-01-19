package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
        "os"
	"runtime/debug"
	"strings"
)

var (
	greeting = flag.String("g", "Hello", "Greet with `greeting`")
	addr     = flag.String("addr", "0.0.0.0:8080", "address to serve")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: service [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// ... (unchanged main function)

	http.HandleFunc("/", greet)
	http.HandleFunc("/version", version)
	http.HandleFunc("/user-form", userForm)
	http.HandleFunc("/process-form", processForm)
        http.HandleFunc("/image-page", imagePage)

	log.Printf("Server is listening on http://%s\n", *addr)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func userForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html>\n<html><head><style>")
	fmt.Fprintf(w, "body { font-family: 'Arial', sans-serif; text-align: center; margin: 50px; }")
	fmt.Fprintf(w, "h1 { color: #333; }")
	fmt.Fprintf(w, "form { display: inline-block; text-align: left; }")
	fmt.Fprintf(w, "</style></head><body>")
	fmt.Fprintf(w, "<h1>%s, enter your information:</h1>\n", *greeting)
	fmt.Fprintf(w, "<form action='/process-form' method='post'>")
	fmt.Fprintf(w, "Name: <input type='text' name='name' required><br>")
	fmt.Fprintf(w, "Email: <input type='email' name='email' required><br>")
	fmt.Fprintf(w, "Favorite Cloud: <select name='favoriteCloud'>")
	fmt.Fprintf(w, "<option value='Cirrus'>Cirrus</option>")
	fmt.Fprintf(w, "<option value='Altocumulus'>Altocumulus</option>")
	fmt.Fprintf(w, "<option value='Cumulus'>Cumulus</option>")
	fmt.Fprintf(w, "<option value='Cumulonimbus'>Cumulonimbus</option>")
	fmt.Fprintf(w, "<option value='Stratus'>Stratus</option>")
	fmt.Fprintf(w, "</select><br>")
	fmt.Fprintf(w, "<input type='submit' value='Submit'>")
	fmt.Fprintf(w, "</form>")
	fmt.Fprintf(w, "</body></html>")
}

func processForm(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	// Extract form values
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	favoriteCloud := r.Form.Get("favoriteCloud")

	// Save the form data to the database (simulated for now)
	saveFormDataToDatabase(name, email, favoriteCloud)

	// Redirect to the page with the image
	http.Redirect(w, r, "/image-page", http.StatusSeeOther)
}

func saveFormDataToDatabase(name, email, favoriteCloud string) {
	// Simulated database operation
	// Replace this with actual database interaction
	fmt.Printf("Saving to database: Name=%s, Email=%s, Favorite Cloud=%s\n", name, email, favoriteCloud)
}

func imagePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html>\n<html><head><style>")
	fmt.Fprintf(w, "body { font-family: 'Arial', sans-serif; text-align: center; margin: 50px; }")
	fmt.Fprintf(w, "h1 { color: #333; }")
	fmt.Fprintf(w, "</style></head><body>")
	fmt.Fprintf(w, "<h1>Thank you for submitting your information!</h1>\n")
	fmt.Fprintf(w, "<p>Look at you, You've said it all...</p>")
	fmt.Fprintf(w, "<img src='/your-image-url' alt='Welcome Image'>")
	fmt.Fprintf(w, "</body></html>")
}
func greet(w http.ResponseWriter, r *http.Request) {
	        // Placeholder
}

func version(w http.ResponseWriter, r *http.Request) {
		        // Placeholder
}

