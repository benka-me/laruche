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

func (namespace Namespace) GetAuthor() string {
	author, _, err := Explode(string(namespace))
	if err != nil {
		return ""
	}
	return author
}
func (namespace Namespace) GetName() string {
	_, name, err := Explode(string(namespace))
	if err != nil {
		return ""
	}
	return name
}
func implode(author, name string) Namespace {
	return Namespace(fmt.Sprintf("%s/%s", author, name))
}

func (np Namespace) String() string {
	return string(np)
}

func (nps Namespaces) String() []string {
	arr := make([]string, len(nps))
	for i, url := range nps {
		arr[i] = string(url)
	}
	return arr
}

func appendOne(dest Namespaces, src Namespace) Namespaces {
	if !dest.Contains(src) {
		return append(dest, src)
	}
	return dest
}
func appendOneString(dest []string, src string) []string {
	if !arrayContainsString(dest, src) {
		return append(dest, src)
	}
	return dest
}

func (nps *Namespaces) Append(src ...Namespace) *Namespaces {
	*nps = AppendUnique(*nps, src...)
	return nps
}

func AppendUniqueString(dest []string, src ...string) []string {
	if dest == nil {
		dest = make([]string, 0)
	}
	if len(src) == 0 {
		return dest
	} else if len(src) == 1 {
		return appendOneString(dest, src[0])
	} else {
		return AppendUniqueString(appendOneString(dest, src[0]), src[1:]...)
	}
}

func AppendUnique(dest Namespaces, src ...Namespace) Namespaces {
	if dest == nil {
		dest = make(Namespaces, 0)
	}
	if len(src) == 0 {
		return dest
	} else if len(src) == 1 {
		return appendOne(dest, src[0])
	} else {
		return AppendUnique(appendOne(dest, src[0]), src[1:]...)
	}
}

func arrayContainsString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
func (nps Namespaces) Contains(str Namespace) bool {
	for _, a := range nps {
		if a == str {
			return true
		}
	}
	return false
}

func (nps *Namespaces) PushUnique(new Namespace) {
	for _, a := range *nps {
		if a == new {
			return
		}
	}
	*nps = AppendUnique(*nps, new)
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
