// curl --http2 --insecure https://localhost:8080
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
}
