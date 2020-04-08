package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

const (
	//localEtc     = "/usr/local/etc/lar"
	localVar     = "/usr/local/var/lar"
	databasePath = localVar + "/laruche.db"
)

var (
	LaruchePath = os.Getenv("HOME") + "/laruche"
	SourcePath  = os.Getenv("GOPATH") + "/src"
	db, dberr   = gorm.Open("sqlite3", databasePath)
)

type State struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	AuthToken string
}

func initDB() {
	if dberr != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&State{})
	db.AutoMigrate(&Bee{})

	// sheetcheat
	// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	//var product Product
	//db.First(&product, 1) // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&product)
}

// Init function will migrate database if it's not ready
func Init() *State {
	initDB()
	state := &State{}
	return state.GetCredential()
}
