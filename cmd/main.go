package main

import(
 "core-banking/internal/fondation"
 	"os"
)
func main() {
	cfg := fondation.Load()
	// db := fondation.NewDB(cfg.DB)
	// defer db.Close()


	serverAddress := ":" + cfg.App.Port
	fmt.Println("Server Go berjalan di port", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}