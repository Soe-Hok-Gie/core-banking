package fondation
import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DatabaseConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() *Config{
	// err:= godotenv.Load()
	// if err != nil{
	// 	log.Fatal(".env error")
	// }
	// Baris ini akan membaca .env jika ada, namun jika tidak ada (seperti di GitHub) kode akan lanjut terus tanpa eror
_ = godotenv.Load() 


	cfg:=&Config{
		App:AppConfig{
			Name :os.Getenv ("APP_NAME"),
			Port :os.Getenv("APP_PORT"),
			Env :os.Getenv("APP_ENV"),
		},
		DB:DatabaseConfig{
			Host:os.Getenv("DB_HOST"),
			Port:os.Getenv("DB_PORT"),
			User:os.Getenv("DB_USER"),
			Password:os.Getenv("DB_PASSWORD"),
			Name:os.Getenv("DB_NAME"),
		},
	}
return cfg
}