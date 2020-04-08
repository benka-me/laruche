package laruche

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/go-playground/validator"
	"strconv"
	"strings"
)

func InitBeeAskUser() *Bee {
	bee := &Bee{}
	scan.V = validator.New()

	bee.Name = strings.ToLower(scan.Step(
		"Name of the new bee micro-service ",
		"required,lte=20,gte=3",
		scan.IsAlphanumDash))

	bee.PkgName = strings.ToLower(scan.Step(
		"Package name (2 - 7 chars long, shorter is better) for packages and types building",
		"required,lte=7,gte=2,alpha",
		func(s string) error { return nil }))

	bee.Repo = strings.Replace(scan.Step(
		"Git repository",
		"required,gte=5",
		func(s string) error { return nil }), " ", "", -1)

	pInt, _ := strconv.Atoi(scan.Step(
		"Port",
		"required,number",
		func(s string) error { return nil }))
	bee.Port = int32(pInt)

	bee.FillDefaultMeta()

	return bee
}

func (bee *Bee) FillDefaultMeta() {
	bee.PkgNameCamel = strings.Title(scan.KebabToCamelCase(bee.PkgName))
	bee.Author = conf.GetUsername()
	bee.ProtoSetup = &ProtoSetup{
		Files: []string{
			fmt.Sprintf("%s.proto", bee.PkgName),
			fmt.Sprintf("rpc-%s.proto", bee.PkgName),
		},
	}

	bee.Languages = &Languages{
		Go: &Go{Setup: &LanguageSetup{
			Active:       true,
			ProtocBinary: "gogoslick",
		}},
		Javascript: &Javascript{Setup: &LanguageSetup{
			Active: true,
		}},
	}
}
