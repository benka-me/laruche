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
}

// Init function will migrate database if it's not ready
func Init() *State {
	initDB()
	return GetState()
}
