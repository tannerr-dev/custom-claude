
can you write me a go web server that writes to the response handler?
**1746245206082**

---
# Simple Go Web Server

Here's a Go web server that writes a response to the client:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/custom", customHandler)
	
	// Set up server port
	port := ":8080"
	
	// Start the server
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
```

## How to Run This Server

1. Save the code to a file named `server.go`
2. Run it using: `go run server.go`
3. Open a web browser and navigate to `http://localhost:8080` 
4. Try the custom endpoint at `http://localhost:8080/custom?name=YourName`

This example demonstrates:
- Setting response headers
- Writing status codes
- Writing HTML content to the response body
- Handling different URL paths
- Processing query parameters

The server will run until you stop it with Ctrl+C in the terminal.
---

