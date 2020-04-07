package generator

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
)

type Go laruche.Go

func (g Go) ClientsFile(bee *laruche.Bee) error {
	repo := bee.Repo
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)

	//generate clients
	err := Code{
		Interface: bee,
		Target:    fmt.Sprintf("%s/go-pkg/http/rpc/clients.go", repoPath),
		Name:      "clients",
	}.GenerateClientsGo()
	if err != nil {
		return err
	}

	return nil
}

func (g Go) ServerFiles(bee *laruche.Bee) error {
	repo := bee.Repo
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)

	//generate main.go
	main := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/main-go", GoTemplates),
		Target:    fmt.Sprintf("%s/src/%s/main.go", gopath, repo),
		Name:      "main",
	}
	err := main.Generate()
	if err != nil {
		return err
	}

	//generate server-grpc_2.0.go
	server := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/server-grpc-2.0-go", GoTemplates),
		Target:    fmt.Sprintf("%s/src/%s/go-pkg/http/rpc/server-grpc-2.0.go", gopath, repo),
		Name:      "server",
	}
	err = server.Generate()
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
	err = hello.Generate()
	if err != nil {
		return err
	}
	return nil
}

func (g Go) Protoc(bee *laruche.Bee) {
	repoPath := fmt.Sprintf("%s/src/%s", gopath, bee.Repo)
	goOut := fmt.Sprintf("%s/src", gopath)
	args := make([]string, 3)
	args = []string{
		fmt.Sprintf("--proto_path=%s/protobuf", repoPath),
		fmt.Sprintf("-I=%s/src", gopath),
		fmt.Sprintf("--%s_out=plugins=grpc:%s", g.Setup.ProtocBinary, goOut),
	}

	for _, f := range bee.ProtoSetup.Files {
		args = append(args, fmt.Sprintf("%s/protobuf/%s", repoPath, f))
	}
	runProtocCommand(args)
}
