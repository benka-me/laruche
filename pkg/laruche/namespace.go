package laruche

import (
	"errors"
	"fmt"
	"strings"
)

type Namespace string
type Namespaces []Namespace
type Mapped map[Namespace]*bool
type Deps []string

func implode(author, name string) Namespace {
	return Namespace(fmt.Sprintf("%s/%s", author, name))
}

func (np Namespace) NamespaceStr() string {
	return string(np)
}

func (nps Namespaces) Array() []string {
	arr := make([]string, len(nps))
	for i, url := range nps {
		arr[i] = string(url)
	}
	return arr
}

func appendOne(dest Namespaces, src Namespace) Namespaces {
	if !dest.contains(src) {
		return append(dest, src)
	}
	return dest
}

func AppendNamespace(dest Namespaces, src ...Namespace) Namespaces {
	if len(src) == 0 {
		return dest
	} else if len(src) == 1 {
		return appendOne(dest, src[0])
	} else {
		return AppendNamespace(appendOne(dest, src[0]), src[1:]...)
	}
}

func (nps Namespaces) contains(str Namespace) bool {
	for _, a := range nps {
		if a == str {
			return true
		}
	}
	return false
}

func BeesToNamespacesFrom(arr []*Bee) Namespaces {
	namespaces := make(Namespaces, len(arr))
	for i, bee := range arr {
		namespaces[i] = bee.GetNamespace()
	}
	return namespaces
}

func ArrayToNamespaces(arr []string) (Namespaces, error) {
	namespaces := make(Namespaces, len(arr))
	for i, url := range arr {
		_, _, err := Explode(url)
		if err != nil {
			return nil, errors.New(url + " bad namespace")
		}
		namespaces[i] = Namespace(url)
	}
	return namespaces, nil
}

func Explode(id string) (author, name string, err error) {
	arr := strings.Split(id, "/")
	if len(arr) == 2 {
		return arr[0], arr[1], nil
	}
	return "error", "error", errors.New("bad namespace")
}

func (bee *Bee) GetNamespace() Namespace {
	return implode(bee.Author, bee.Name)
}
func (hive *Hive) GetNamespace() Namespace {
	return implode(hive.Author, hive.Name)
}
func (hive *Hive) GetNamespaceStr() string {
	return string(hive.GetNamespace())
}
func (bee *Bee) GetNamespaceStr() string {
	return string(bee.GetNamespace())
}
