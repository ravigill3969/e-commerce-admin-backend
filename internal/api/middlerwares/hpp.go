package middlerwares

import (
	"fmt"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckQuery              bool
	CheckBody               bool
	CheckBodyForContentType string
	Whitelist               []string
}

func Hpp(options HPPOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyForContentType) {
				//filter the body params
				filterBodyParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whiteList []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0]) // first value
			//r.Form.Set(k, v[len(v) - 1])//last value
		}

		if !isWhiteListed(k, whiteList) {
			delete(r.Form, k)
		}
	}

}

func isWhiteListed(param string, whiteList []string) bool {
	for _, v := range whiteList {
		if param == v {
			return true
		}
	}
	return false
}
