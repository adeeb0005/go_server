package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./static/form.html")
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "POST request successful\n")
		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Address: %s\n", address)
		return
	}

	// If not GET or POST
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Serve files from the "static" directory at root level (http://localhost:8000/form.html)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Form POST handler
	http.HandleFunc("/form", formHandler)

	// Hello route
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
