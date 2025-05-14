package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	port := 8080

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logReqDeatils(r)
		fmt.Fprintf(w, "Handling incoming orders\n")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling  users")
	})

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)

	if err != nil {
		fmt.Println("error in starting server....", err)
	}
}

func logReqDeatils(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("received req with http vesion", httpVersion)

	if r.TLS != nil {
		tlsVesrion := getTLSVersionName(r.TLS.Version)
		fmt.Println("received req with tls vesion", tlsVesrion)

	} else {
		fmt.Println("received req without tls")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "tls 1.0"

	case tls.VersionTLS11:
		return "tls 1.1"

	case tls.VersionTLS12:
		return "tls 1.2"

	case tls.VersionTLS13:
		return "tls 1.3"

	default:
		return "Unknown TLS version"
	}

}
