package generator

import (
	"fmt"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"os"
)

var dirPerm os.FileMode = 0755

func mkdirAll(s string) error {
	return os.MkdirAll(s, dirPerm)
}

func agnosticServerFiles(bee *laruche.Bee) error {
	repo := bee.Repo
	pkgName := bee.PkgName
	repoPath := fmt.Sprintf("%s/%s", sourcePath, repo)

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
		if err := mkdirAll(dir); err != nil {
			return err
		}
	}

	//generate defs.proto file
	err := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/defs.proto", ProtobufTemplates),
		Target:    fmt.Sprintf("%s/protobuf/%s.proto", repoPath, pkgName),
		Name:      "defs",
	}.generate()
	if err != nil {
		return err
	}

	//generate rpc.proto file
	err = Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/rpc.proto", ProtobufTemplates),
		Target:    fmt.Sprintf("%s/protobuf/rpc-%s.proto", repoPath, pkgName),
		Name:      "rpc",
	}.generate()
	if err != nil {
		return err
	}
	return nil
}
