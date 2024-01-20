package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"
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
	flag.Usage = usage
	flag.Parse()

	http.HandleFunc("/", greet)
	http.HandleFunc("/version", version)

	log.Printf("Server is listening on http://%s\n", *addr)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n<title>Version Information</title>\n<style>body {font-family: 'Arial', sans-serif;}</style>\n</head>\n<body>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
	fmt.Fprintf(w, "</pre>\n</body>\n</html>")
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Guest"
	}

	greetMessage := getDynamicGreeting()

	fmt.Fprintf(w, "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n<title>Personalized Greeting</title>\n<style>body {font-family: 'Arial', sans-serif; text-align: center; padding: 50px;}</style>\n</head>\n<body>\n")
	fmt.Fprintf(w, "<h1>%s, %s!</h1>\n", greetMessage, html.EscapeString(name))
	fmt.Fprintf(w, "<p>Feel free to explore and make yourself at home.</p>\n</body>\n</html>")
}

func getDynamicGreeting() string {
	currentHour := time.Now().Hour()

	switch {
	case currentHour >= 5 && currentHour < 12:
		return "Good morning"
	case currentHour >= 12 && currentHour < 17:
		return "Good afternoon"
	default:
		return "Good evening"
	}
}

