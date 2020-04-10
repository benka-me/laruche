package form

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/cli/scan"
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/go-playground/validator"
	"strconv"
	"strings"
)

func InitGatewayAskUser() *laruche.Bee {
	bee := &laruche.Bee{}
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

	bee.Port = int32(8080)
	FillDefaultMeta(bee)
	return bee
}

func InitClientAskUser() *laruche.Bee {
	bee := &laruche.Bee{}
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
	FillDefaultMeta(bee)
	return bee
}

func InitServiceAskUser() *laruche.Bee {
	bee := &laruche.Bee{}
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

	FillDefaultMeta(bee)

	return bee
}

func FillDefaultMeta(bee *laruche.Bee) {
	bee.PkgNameCamel = strings.Title(scan.KebabToCamelCase(bee.PkgName))
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
