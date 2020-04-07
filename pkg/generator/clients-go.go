package generator

import (
	"bytes"
	"fmt"
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/laruche/pkg/laruche"
	"io/ioutil"
	"os"
	"text/template"
)

type PkgNames struct {
	PkgName      string
	PkgNameCamel string
}

type Clients struct {
	Import string
	Vars   string
	Init   string
	Return string
}

func (code Code) Import() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-import-go", GoTemplates))
	_if.ErrorExit("generate clients: read file import", err)

	_import, err := template.New("import").Parse(string(dat))
	_if.ErrorExit("generate clients: import parse template", err)

	buf := bytes.Buffer{}
	err = _import.Execute(&buf, code.Interface)
	_if.ErrorExit("generate clients: import execute", err)

	return buf.String()
}

func (code Code) Vars() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-vars-go", GoTemplates))
	_if.ErrorExit("generate clients: vars read file", err)

	_vars, err := template.New("vars").Parse(string(dat))
	_if.ErrorExit("generate clients: vars parse", err)

	buf := bytes.Buffer{}
	err = _vars.Execute(&buf, code.Interface)
	_if.ErrorExit("generate clients: vars execute", err)

	return buf.String()
}
func (code Code) Init() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-init-go", GoTemplates))
	_if.ErrorExit("generate clients: init read file", err)

	_init, err := template.New("init").Parse(string(dat))

	_if.ErrorExit("generate clients: init parse", err)

	buf := bytes.Buffer{}
	err = _init.Execute(&buf, code.Interface)
	_if.ErrorExit("generate clients: init execute", err)

	return buf.String()
}

func (code Code) Return() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-return-go", GoTemplates))
	_if.ErrorExit("generate clients: return readfile", err)

	_return, err := template.New("return").Parse(string(dat))
	_if.ErrorExit("generate clients: return parse", err)

	buf := bytes.Buffer{}
	err = _return.Execute(&buf, code.Interface)
	_if.ErrorExit("generate clients: return execute", err)

	return buf.String()
}

func (code Code) GenerateClientsGo() error {
	cl := Clients{}
	clientsPath := fmt.Sprintf("%s/clients-go", GoTemplates)
	dat, err := ioutil.ReadFile(clientsPath)
	if err != nil {
		return err
	}

	//bee := code.Interface.(*laruche.Bee)

	//deps, err := request.GetRemoteBees(bee.GetSubDependencies())
	//if err != nil {
	//	fmt.Println("!-- to generate clients please connect to hiveof.services with login cmd --!")
	//}
	deps := []laruche.Bee{}

	for _, dep := range deps {
		co := Code{
			Interface: dep,
			Name:      dep.Name,
		}
		cl = Clients{
			fmt.Sprintf("%s%s", cl.Import, co.Import()),
			fmt.Sprintf("%s%s", cl.Vars, co.Vars()),
			fmt.Sprintf("%s%s", cl.Init, co.Init()),
			fmt.Sprintf("%s%s", cl.Return, co.Return()),
		}
	}

	tmpl, err := template.New("clients").Parse(string(dat))
	if err != nil {
		return err
	}

	f, err := os.Create(code.Target)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, cl)
	if err != nil {
		return err
	}
	return nil
}
