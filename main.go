package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

func main() {
	// Create a request limiter per IP
	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: 1 * 60 * 60 * 1000000000})

	http.Handle("/", tollbooth.LimitFuncHandler(lmt, func(w http.ResponseWriter, r *http.Request) {
		// Set some security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")

		fmt.Fprintf(w, "Hello, from the other siiiiiideeeee!!")
	}))

	http.Handle("/test", tollbooth.LimitFuncHandler(lmt, func(w http.ResponseWriter, r *http.Request) {
		// Set some security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")

		fmt.Fprintf(w, "This is a GET request for /test")
	}))

	// Get the port number from the environment variable
	port := os.Getenv("API_PORT")
	if port == "" {
		log.Fatal("API_PORT environment variable not set")
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
