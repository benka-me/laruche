package generator

import (
	"fmt"
	form2 "github.com/benka-me/laruche/go-pkg/cli/form"
	config2 "github.com/benka-me/laruche/go-pkg/config"
	laruche2 "github.com/benka-me/laruche/go-pkg/laruche"
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

var sourcePath = config2.SourcePath
var Templates = fmt.Sprintf("%s/github.com/benka-me/laruche/go-pkg/generator/templates", sourcePath)
var GoTemplates = fmt.Sprintf("%s/go", Templates)
var ProtobufTemplates = fmt.Sprintf("%s/protobuf", Templates)

func GenerateMainService(bee *laruche2.Bee) error {
	form2.FillDefaultMeta(bee)
	err := agnosticServerFiles(bee)
	if err != nil {
		return err
	}

	generators, err := GetLangs(bee.Languages)
	if err != nil {
		return err
	}

	for _, lang := range *generators {
		lang.Protoc(bee)
	}

	err = GetDevLang(bee).MainServer(bee)
	if err != nil {
		return err
	}

	err = GetDevLang(bee).ClientsFile(bee)
	if err != nil {
		return err
	}
	return nil
}

func GenerateMainClient(bee *laruche2.Bee) error {
	form2.FillDefaultMeta(bee)

	devLang := GetDevLang(bee)

	err := devLang.MainClient(bee)
	if err != nil {
		return err
	}

	err = GetDevLang(bee).ClientsFile(bee)
	if err != nil {
		return err
	}
	return nil
}

func GenerateClients(bee *laruche2.Bee) error {
	err := GetDevLang(bee).ClientsFile(bee)
	return err
}

func (code Code) generate() error {
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
