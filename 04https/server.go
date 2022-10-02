package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	// SSL証明書cert.pem/秘密鍵key.pemを用いてhttpsによるサービスを開始
	// cert.pem/key.pemの作成コード: 05gencert/gencert.go
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
