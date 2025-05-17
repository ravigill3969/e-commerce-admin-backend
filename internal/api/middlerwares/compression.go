package middlerwares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check if client accept gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}

		//set the res header
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}

		next.ServeHTTP(w, r)
		fmt.Println("compressing this garbage")
	})
}

//gzip response writer will wrap the http .REsponsewriter to write  gzipped responses

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
