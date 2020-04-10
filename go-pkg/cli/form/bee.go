package form

import (
	"fmt"
	scan2 "github.com/benka-me/laruche/go-pkg/cli/scan"
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/go-playground/validator"
	"strconv"
	"strings"
)

func InitGatewayAskUser() *laruche.Bee {
	bee := &laruche.Bee{}
	scan2.V = validator.New()

	bee.IsGateway = true
	bee.Name = strings.ToLower(scan2.Step(
		"Name of the new bee micro-service ",
		"required,lte=20,gte=3",
		scan2.IsAlphanumDash))

	bee.PkgName = strings.ToLower(scan2.Step(
		"Package name (2 - 7 chars long, shorter is better) for packages and types building",
		"required,lte=7,gte=2,alpha",
		func(s string) error { return nil }))

	bee.Repo = strings.Replace(scan2.Step(
		"Git repository",
		"required,gte=5",
		func(s string) error { return nil }), " ", "", -1)

	bee.Port = int32(8080)
	FillDefaultMeta(bee)
	return bee
}

func InitClientAskUser() *laruche.Bee {
	bee := &laruche.Bee{}
	scan2.V = validator.New()

	bee.Name = strings.ToLower(scan2.Step(
		"Name of your client",
		"required,lte=20,gte=3",
		scan2.IsAlphanumDash))

	//bee.PkgName = strings.ToLower(scan.Step(
	//	"Package name (2 - 7 chars long, shorter is better) for packages and types building",
	//	"required,lte=7,gte=2,alpha",
	//	func(s string) error { return nil }))
	//
	bee.Repo = strings.Replace(scan2.Step(
		"Git repository",
		"required,gte=5",
		func(s string) error { return nil }), " ", "", -1)
	FillDefaultMeta(bee)
	return bee
}

func InitServiceAskUser() *laruche.Bee {
	bee := &laruche.Bee{}
	scan2.V = validator.New()

	bee.Name = strings.ToLower(scan2.Step(
		"Name of the new bee micro-service ",
		"required,lte=20,gte=3",
		scan2.IsAlphanumDash))

	bee.PkgName = strings.ToLower(scan2.Step(
		"Package name (2 - 7 chars long, shorter is better) for packages and types building",
		"required,lte=7,gte=2,alpha",
		func(s string) error { return nil }))

	bee.Repo = strings.Replace(scan2.Step(
		"Git repository",
		"required,gte=5",
		func(s string) error { return nil }), " ", "", -1)

	pInt, _ := strconv.Atoi(scan2.Step(
		"Port",
		"required,number",
		func(s string) error { return nil }))
	bee.Port = int32(pInt)

	FillDefaultMeta(bee)

	return bee
}

func FillDefaultMeta(bee *laruche.Bee) {
	bee.PkgNameCamel = strings.Title(scan2.KebabToCamelCase(bee.PkgName))
	bee.Author = config.GetState().Username
	bee.ProtoSetup = &laruche.ProtoSetup{
		Files: []string{
			fmt.Sprintf("%s.proto", bee.PkgName),
			fmt.Sprintf("rpc-%s.proto", bee.PkgName),
		},
	}

	bee.Languages = &laruche.Languages{
		Go: &laruche.Go{Setup: &laruche.LanguageSetup{
			Active:       true,
			ProtocBinary: "gogoslick",
		}},
		Javascript: &laruche.Javascript{Setup: &laruche.LanguageSetup{
			Active: true,
		}},
	}
}
