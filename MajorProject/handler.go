package MajorProject

import (
	"ecommerce/dbc"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

type Handler struct {
	DB *gorm.DB
}

func Dbhand() Handler {
	err := godotenv.Load(".env")
	if err != nil {
		return Handler{}
	}
	b := os.Getenv("dns")
	// a := os.Getenv("dns")

	return Handler{
		DB: dbc.Dbinit(b),
	}
}
