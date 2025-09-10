package main

import (
	"log"
	"net/http"
	"real-camera/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/ws", handlers.WSHandler)
	
	log.Println("サーバーを開始します: http://localhost:8080")
	log.Println("OSCターゲット: 127.0.0.1:8765")
	log.Fatal(http.ListenAndServe(":8080", nil))
}