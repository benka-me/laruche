package config

import (
	"github.com/benka-me/laruche/pkg/cli/scan"
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
		if !scan.StepBool("Namespace existing on your local machine, are you sur you want to re-generate files?") {
			os.Exit(9)
		}
	}
	RemoveBee(new.GetNamespace())
	ret := db.Create(&Bee{
		ID:   new.GetNamespaceStr(),
		Path: new.Repo,
	})
	return ret.Value.(*Bee)
}

func RemoveBee(namespace laruche.Namespace) *Bee {
	b := &Bee{}
	b = db.Delete(b, "id = ?", namespace).Value.(*Bee)
	return b
}
func GetBee(namespace laruche.Namespace) *Bee {
	b := &Bee{}
	b = db.Find(b, "id = ?", namespace).Value.(*Bee)
	return b
}

func (b *Bee) GetPath() string {
	return b.Path
}
