package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.tailwind/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Beranda)
	http.HandleFunc("/profil", handlers.ProfilHandler)
	http.HandleFunc("/ekstrakulikuler", handlers.EkstrakulikulerHandler)
	http.HandleFunc("/galeri", handlers.GaleriHandler)


	fmt.Println("Server Ready Boss: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}