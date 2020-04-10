package generator

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
)

type Javascript laruche.Javascript

func (js Javascript) ClientsFile(bee *laruche.Bee) error {
	return nil
}

func (js Javascript) ServerFiles(bee *laruche.Bee) error {
	return nil
}

func (js Javascript) Protoc(bee *laruche.Bee) {
	repoPath := fmt.Sprintf("%s/%s", sourcePath, bee.Repo)
	//bin := fmt.Sprintf("%s/github.com/benka-me/hive/js-pkg/node_modules/.bin/protoc-gen-ts", sourcePath)
	bin := "."
	jsOut := fmt.Sprintf("%s/js-pkg/src/protobuf", repoPath)
	args := make([]string, 5)
	args = []string{
		fmt.Sprintf("--proto_path=%s/protobuf", repoPath),
		fmt.Sprintf("-I=%s", sourcePath),
		fmt.Sprintf("--plugin=protoc-gen-ts=%s", bin),
		//fmt.Sprintf("--ts_out=service=true:%s", jsOut),
		fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", jsOut),
	}

	for _, f := range bee.ProtoSetup.Files {
		args = append(args, fmt.Sprintf("%s", f))
	}
	runProtocCommand(args)
}
