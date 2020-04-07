package laruche

import (
	"fmt"
	"strings"
)

type Namespace string
type Namespaces []Namespace
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
func ArrayToNamespaces(arr []string) Namespaces {
	namespaces := make(Namespaces, len(arr))
	for i, url := range arr {
		namespaces[i] = Namespace(url)
	}
	return namespaces
}

func Explode(id string) (author, name string) {
	arr := strings.Split(id, "/")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return "error", "error"
}

func (b *Bee) GetNamespace() Namespace {
	return implode(b.Author, b.Name)
}
func (h *Hive) Namespace() Namespace {
	return implode(h.Author, h.Name)
}
func (h *Hive) NamespaceStr() string {
	return string(h.Namespace())
}

func (b *Bee) GetNamespaceStr() string {
	return string(b.GetNamespace())
}
