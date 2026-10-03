package main

import (
	"core-banking/internal/fondation"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := fondation.Load()
	// db := fondation.NewDB(cfg.DB)
	// defer db.Close()

	serverAddress := ":" + cfg.App.Port
	fmt.Println("Server Go berjalan di port", serverAddress)

	// Membuat router/mux kosong karena variabel 'r' belum didefinisikan sebelumnya
	r := http.NewServeMux()

	// Menambahkan contoh endpoint agar server bisa diakses
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Core Banking API is running")
	})

	log.Fatal(http.ListenAndServe(serverAddress, r))
}
