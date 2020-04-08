package config

import (
	"github.com/jinzhu/gorm"
	"os"
)

const (
	localEtc = "/usr/local/etc/lar"
	localVar = "/usr/local/var/lar"
)

var (
	laruchePath = os.Getenv("HOME") + "/laruche"
	sourcePath  = os.Getenv("GOPATH") + "/src"
)

type State struct {
	gorm.Model
	Username  string
	AuthToken string
}

func Init() State {

	return State{}
}
