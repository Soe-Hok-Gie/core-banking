package main

import (
	"core-banking/internal/core/user/controller"
	"core-banking/internal/core/user/repository"
	"core-banking/internal/core/user/service"
	"core-banking/internal/fondation"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := fondation.Load()
	db := fondation.NewDB(cfg.DB)
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(authService, userService)

	r := mux.NewRouter()
	r.HandleFunc("/auth/register", userController.Register).Methods("POST")
	r.HandleFunc("/auth/login", userController.Login).Methods("POST")

	// serverAddress := ":" + cfg.App.Port
	// fmt.Println("Server Go berjalan di port", serverAddress)

	// // Menambahkan contoh endpoint agar server bisa diakses
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Core Banking API is running")
	// })

	// log.Fatal(http.ListenAndServe(serverAddress, r))

	fmt.Println("Database Connected")
	serverAddress := ":" + cfg.App.Port
	fmt.Println("server go berjalan", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
