package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load() // Memuat file .env

	// Ambil DATABASE_URL langsung dari environment
	dsn := os.Getenv("https://postgres:ZNUffYdAUxjLBAiZTgJHrPuJNwwsUvOS@viaduct.proxy.rlwy.net:26997/railway")
	if dsn == "" {
		panic("DATABASE_URL tidak ditemukan di environment variables")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi ke database: " + err.Error())
	}

	fmt.Println("Database berhasil terkoneksi!")
}
