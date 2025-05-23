package main

import (
	mw "admin-backend/internal/api/middlerwares"
	"admin-backend/internal/api/repository/sqlconnect"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	// "time"

	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root endpoint\n")
}
func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET: Returning teacher info")

	case http.MethodPost:
		fmt.Fprintln(w, "POST: Creating a new teacher")

	case http.MethodPut:
		fmt.Fprintln(w, "PUT: Replacing teacher details")

	case http.MethodPatch:
		fmt.Fprintln(w, "PATCH: Updating partial teacher data")

	case http.MethodDelete:
		fmt.Fprintln(w, "DELETE: Removing teacher record")

	case http.MethodOptions:
		// Inform client of supported methods
		w.Header().Set("Allow", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	err := godotenv.Load()

	if err != nil {
		return
	}
	_, err = sqlconnect.ConnectDB()

	if err != nil {
		fmt.Println("Unable to connect", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/user", teachersHandler)

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:              true,
	// 	CheckBody:               true,
	// 	CheckBodyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:               []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middlerware(mw.ResponseTimeMiddlerware(mw.Compression(mw.Hpp(hppOptions)(mw.SecurityHeader(mux))))))
	// secureMux := applyMiddleware(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeader, mw.ResponseTimeMiddlerware, rl.Middlerware, mw.Cors)
	port := os.Getenv("API_PORT")
	secureMux := mw.SecurityHeader(mux)
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

// Middleware is a fn  that wraps on http.handler with addiotanial fntionality
type Middlerware func(http.Handler) http.Handler

// func applyMiddleware(handler http.Handler, middlerwares ...Middlerware) http.Handler {
// 	for _, middlerware := range middlerwares {
// 		handler = middlerware(handler)
// 	}
// 	return handler
// }
