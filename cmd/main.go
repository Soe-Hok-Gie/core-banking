package main

import (
	"core-banking/internal/core/user/controller"
	"core-banking/internal/core/user/repository"
	"core-banking/internal/core/user/service"
	"core-banking/internal/fondation"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := fondation.Load()
	db := fondation.NewDB(cfg.DB)
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	AuthService := service.NewAuthService(userRepository)
	userController := controller.NewUserController(AuthService)

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
