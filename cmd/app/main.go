package main

import (
	"log"
	"net/http"
)

func main() {
	// Простейший обработчик для проверки
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong 🚀"))
	})

	log.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
