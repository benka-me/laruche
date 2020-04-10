package generator

import (
	"bytes"
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
	"io/ioutil"
	"os"
	"text/template"
)

type Go laruche.Go
type GoClients struct {
	Import string
	Vars   string
	Init   string
	Return string
}

func (g Go) MainServer(bee *laruche.Bee) error {
	repo := bee.Repo
	repoPath := fmt.Sprintf("%s/%s", sourcePath, repo)

	//generate main.go
	main := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/main-go", GoTemplates),
		Target:    fmt.Sprintf("%s/%s/main.go", sourcePath, repo),
		Name:      "main",
	}
	err := main.generate()
	if err != nil {
		return err
	}

	//generate server-grpc_2.0.go
	server := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/server-grpc-2.0-go", GoTemplates),
		Target:    fmt.Sprintf("%s/%s/go-pkg/http/rpc/server-grpc-2.0.go", sourcePath, repo),
		Name:      "server",
	}
	err = server.generate()
	if err != nil {
		return err
	}

	//generate hello-world.go
	hello := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/hello-world-go", GoTemplates),
		Target:    fmt.Sprintf("%s/go-pkg/http/rpc/hello-world.go", repoPath),
		Name:      "hello",
	}
	err = hello.generate()
	if err != nil {
		return err
	}
	return nil
}

func (g Go) MainClient(bee *laruche.Bee) error {
	repo := bee.Repo
	repoPath := fmt.Sprintf("%s/%s", sourcePath, repo)
	err := mkdirAll(fmt.Sprintf("%s/go-pkg/http/rpc", repoPath))
	if err != nil {
		return err
	}
	//generate main.go
	main := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/main-client-go", GoTemplates),
		Target:    fmt.Sprintf("%s/%s/main.go", sourcePath, repo),
		Name:      "main",
	}
	err = main.generate()
	if err != nil {
		return err
	}
	return nil
}

func (g Go) Protoc(bee *laruche.Bee) {
	repoPath := fmt.Sprintf("%s/%s", sourcePath, bee.Repo)
	goOut := fmt.Sprintf("%s", sourcePath)
	args := make([]string, 3)
	args = []string{
		fmt.Sprintf("--proto_path=%s/protobuf", repoPath),
		fmt.Sprintf("-I=%s", sourcePath),
		fmt.Sprintf("--%s_out=plugins=grpc:%s", g.Setup.ProtocBinary, goOut),
	}

	for _, f := range bee.ProtoSetup.Files {
		args = append(args, fmt.Sprintf("%s/protobuf/%s", repoPath, f))
	}
	runProtocCommand(args)
}

func (g Go) ClientsFile(bee *laruche.Bee) error {
	repoPath := fmt.Sprintf("%s/%s", sourcePath, bee.Repo)
	code := Code{
		Interface: bee,
		Target:    fmt.Sprintf("%s/go-pkg/http/rpc/clients.go", repoPath),
		Name:      "clients",
	}

	cl := GoClients{}
	clientsPath := fmt.Sprintf("%s/clients-go", GoTemplates)
	dat, err := ioutil.ReadFile(clientsPath)
	if err != nil {
		return err
	}

	deps, err := absolute.GetBeez(bee.GetDependencies())
	if err != nil {
		fmt.Println("!-- to generate clients please connect to hiveof.services with login cmd --!")
	}

	deps.Map(func(i int, dep *laruche.Bee) error {
		fmt.Println("client for ", dep.GetNamespace())
		co := Code{
			Interface: dep,
			Name:      dep.Name,
		}
		cl = GoClients{
			fmt.Sprintf("%s%s", cl.Import, co.blockImport()),
			fmt.Sprintf("%s%s", cl.Vars, co.blockVars()),
			fmt.Sprintf("%s%s", cl.Init, co.blockInit()),
			fmt.Sprintf("%s%s", cl.Return, co.blockReturn()),
		}
		return nil
	})

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

func (code Code) blockImport() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-import-go", GoTemplates))
	if err != nil {
		fmt.Println("generate clients: read file import", err)
	}

	_import, err := template.New("import").Parse(string(dat))
	if err != nil {
		fmt.Println("generate clients: import parse template", err)
	}

	buf := bytes.Buffer{}
	err = _import.Execute(&buf, code.Interface)
	if err != nil {
		fmt.Println("generate clients: import execute", err)
	}

	return buf.String()
}

func (code Code) blockVars() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-vars-go", GoTemplates))
	if err != nil {
		fmt.Println("generate clients: vars read file", err)
	}

	_vars, err := template.New("vars").Parse(string(dat))
	if err != nil {
		fmt.Println("generate clients: vars parse", err)
	}

	buf := bytes.Buffer{}
	err = _vars.Execute(&buf, code.Interface)
	if err != nil {
		fmt.Println("generate clients: vars execute", err)
	}

	return buf.String()
}

func (code Code) blockInit() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-init-go", GoTemplates))
	if err != nil {
		fmt.Println("generate clients: init read file", err)
	}

	_init, err := template.New("init").Parse(string(dat))
	if err != nil {
		fmt.Println("generate clients: init parse", err)
	}

	buf := bytes.Buffer{}
	err = _init.Execute(&buf, code.Interface)
	if err != nil {
		fmt.Println("generate clients: init execute", err)
	}

	return buf.String()
}

func (code Code) blockReturn() string {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/clients-return-go", GoTemplates))
	if err != nil {
		fmt.Println("generate clients: return readfile", err)
	}

	_return, err := template.New("return").Parse(string(dat))
	if err != nil {
		fmt.Println("generate clients: return parse", err)
	}

	buf := bytes.Buffer{}
	err = _return.Execute(&buf, code.Interface)
	if err != nil {
		fmt.Println("generate clients: return execute", err)
	}

	return buf.String()
}
