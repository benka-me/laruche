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

func AddBee(new *laruche.Bee) error {
	test, _ := GetBee(new.GetNamespace())
	if test != nil {
		if !scan.StepBool("GetNamespace existing on your local machine, are you sur you want to re-generate files?") {
			os.Exit(9)
		}
	}
	_, _ = RemoveBee(new.GetNamespace())
	_ = db.Create(&Bee{
		ID:   new.GetNamespaceStr(),
		Path: new.Repo,
	})
	return nil
}

func RemoveBee(namespace laruche.Namespace) (*Bee, error) {
	b := &Bee{}
	b = db.Delete(b, "id = ?", namespace).Value.(*Bee)
	return b, nil
}
func GetBee(namespace laruche.Namespace) (*Bee, error) {
	b := &Bee{}
	b = db.Find(b, "id = ?", namespace).Value.(*Bee)
	return b, nil
}

func (b *Bee) GetPath() string {
	if b == nil {
		return ""
	}
	return b.Path
}
