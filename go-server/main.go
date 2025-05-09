package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"html/template"
)
// Basic handler function that writes to the response
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Write a header
	w.Header().Set("Content-Type", "text/html")
	
	// Write the response status code
	w.WriteHeader(http.StatusOK)
	
	// Write to the response body
	tmpl := template.Must(template.ParseFiles("home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


// Basic handler function that writes to the response
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write a header
	w.Header().Set("Content-Type", "text/html")
	
	// Write the response status code
	w.WriteHeader(http.StatusOK)
	
	// Write to the response body
	fmt.Fprintf(w, "<h1>Hello from Go Web Server!</h1>")
	fmt.Fprintf(w, "<p>Current time: %s</p>", time.Now().Format(time.RFC1123))
	fmt.Fprintf(w, "<p>Request path: %s</p>", r.URL.Path)
	fmt.Fprintf(w, "<p>Request method: %s</p>", r.Method)
}

// Custom handler with more details
func customHandler(w http.ResponseWriter, r *http.Request) {
	// Set multiple headers
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Custom-Header", "Custom Value")
	
	// Get query parameters
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	
	// Write response
	fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
	fmt.Fprintf(w, "<p>You can pass '?name=YourName' in the URL to customize this greeting.</p>")
}

func main() {
	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/custom", customHandler)
	
	// Set up server port
	port := ":8080"
	
	// Start the server
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
