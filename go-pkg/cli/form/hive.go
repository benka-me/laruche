package form

import (
	"github.com/benka-me/laruche/go-pkg/cli/scan"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/go-playground/validator"
	"strings"
)

func FillMeta(hive *laruche.Hive) {
	hive.PkgNameCamel = strings.Title(scan.KebabToCamelCase(hive.PkgName))
}

func InitHiveAskUser() *laruche.Hive {
	hive := &laruche.Hive{}
	scan.V = validator.New()

	hive.Name = strings.ToLower(scan.Step(
		"Name of the new hive micro-service based application",
		"required,lte=20,gte=3",
		scan.IsAlphanumDash))

	FillMeta(hive)

	return hive
}
