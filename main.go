package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/perqin/shadow-dashboard-go/db"
	"github.com/perqin/shadow-dashboard-go/routers"
)

func main() {
	// Ensure data directory
	if err := os.Mkdir("data", 0755); err != nil {
		panic("Failed to create directory ./data")
	}
	// Prepare SQLite3
	if err := db.OpenDb(); err != nil {
		panic("Failed to open database file")
	}
	defer db.CloseDb()
	// Prepare and startup server
	router := routers.NewRouter()
	router.Run("localhost:4577")
}
