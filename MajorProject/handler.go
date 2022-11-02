package MajorProject

import (
	"ecommerce/dbc"
	"gorm.io/gorm"
	"os"
)

type Handler struct {
	DB *gorm.DB
}

func Dbhand() Handler {
	var DNS = os.Getenv("dns")
	return Handler{
		DB: dbc.Dbinit(DNS),
	}
}
