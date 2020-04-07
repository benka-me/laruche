package generator

import (
	"fmt"
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/laruche/pkg/laruche"
	"os"
)

var dirPerm os.FileMode = 0777

func mkdirAll(s string) {
	err := os.MkdirAll(s, dirPerm)
	_if.ErrorExit("mkdir all "+s, err)
}

func agnosticFiles(bee *laruche.Bee) error {
	repo := bee.Repo
	pkgName := bee.PkgName
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)

	dirs := []string{
		// proto files templates
		fmt.Sprintf("%s/protobuf", repoPath), //proto file dir
		// go
		fmt.Sprintf("%s/go-pkg/http/rpc", repoPath),    //go rpc server dirs
		fmt.Sprintf("%s/go-pkg/%s", repoPath, pkgName), //go protobuf generated package dirs
		// javascript
		fmt.Sprintf("%s/js-pkg/src/protobuf", repoPath), //javascript protobuf generated packages dirs
	}
	for _, dir := range dirs {
		mkdirAll(dir)
	}

	//generate defs.proto file
	err := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/defs.proto", ProtobufTemplates),
		Target:    fmt.Sprintf("%s/protobuf/%s.proto", repoPath, pkgName),
		Name:      "defs",
	}.Generate()
	if err != nil {
		return err
	}

	//generate rpc.proto file
	err = Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/rpc.proto", ProtobufTemplates),
		Target:    fmt.Sprintf("%s/protobuf/rpc-%s.proto", repoPath, pkgName),
		Name:      "rpc",
	}.Generate()
	if err != nil {
		return err
	}
	return nil
}
