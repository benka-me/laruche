package config

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
	"os"
)

type Bee struct {
	ID   string `gorm:"primary_key"`
	Path string
}

func AddBee(new *laruche.Bee) *Bee {
	test := GetBee(new.GetNamespace())
	if test != nil {
		fmt.Println("CONFIG: add bee: namespace already exist")
		os.Exit(0)
	}
	ret := db.Create(&Bee{
		ID:   new.GetNamespaceStr(),
		Path: new.Repo,
	})
	return ret.Value.(*Bee)
}

func GetBee(namespace laruche.Namespace) *Bee {
	b := &Bee{}
	b = db.Find(b, "id = ?", namespace).Value.(*Bee)
	return b
}

func (b *Bee) GetPath() string {
	return b.Path
}
