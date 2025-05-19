package main

import (
	mw "admin-backend/internal/api/middlerwares"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root endpoint\n")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET: Returning user info")
	case http.MethodPost:
		fmt.Println(r.URL.Query())
		fmt.Println(r.URL.Query().Get("name"))

		//parse from data

		err := r.ParseForm()
		if err != nil {
			return
		}

		fmt.Println(r.Form)

		w.Write([]byte("Post method in execution"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	port := 8080

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/user", userHandler)

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   rl.Middlerware(mw.Compression(mw.ResponseTimeMiddlerware(mw.SecurityHeader(mw.Cors(mux))))),
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
