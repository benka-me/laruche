package generator

import (
	"fmt"
	_if "github.com/benka-me/laruche/pkg/if"
	"github.com/benka-me/laruche/pkg/laruche"
	"io/ioutil"
	"os"
	"text/template"
)

type Protobuf struct {
	Repo  string
	Files []string
}

type Code struct {
	Interface interface{}
	Template  string
	Target    string
	Name      string
}

var gopath = os.Getenv("GOPATH")
var Templates = fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/templates", gopath)
var GoTemplates = fmt.Sprintf("%s/go", Templates)
var ProtobufTemplates = fmt.Sprintf("%s/protobuf", Templates)

func GenerateAll(bee *laruche.Bee) {
	bee.FillDefaultMeta()
	err := agnosticFiles(bee)
	_if.ErrorExit("generate all, agnostic files", err)

	generators, err := GetLangs(bee.Languages)
	_if.ErrorExit("generate all, to lang generator", err)

	for _, lang := range *generators {
		lang.Protoc(bee)
	}

	err = GetDevLang(bee).ServerFiles(bee)
	_if.ErrorExit("generate all entry points files", err)

	err = GetDevLang(bee).ClientsFile(bee)
	_if.ErrorExit("46 generate clients file", err)
}

func GenerateClientsFilesFor(bee *laruche.Bee) {
	bee.FillDefaultMeta()

	err := GetDevLang(bee).ClientsFile(bee)
	_if.ErrorExit("generate clients files", err)
}

func (code Code) Generate() error {
	dat, err := ioutil.ReadFile(code.Template)
	if err != nil {
		return err
	}

	tmpl, err := template.New(code.Name).Parse(string(dat))
	if err != nil {
		return err
	}

	f, err := os.Create(code.Target)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, code.Interface)
	if err != nil {
		return err
	}
	return nil
}
