package laruche

import (
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/go-playground/validator"
	"strings"
)

type Hives []*Hive

func (hive *Hive) FillMeta() {
	hive.PkgNameCamel = strings.Title(scan.KebabToCamelCase(hive.PkgName))
}

func InitHiveAskUser() *Hive {
	hive := &Hive{}
	scan.V = validator.New()

	hive.Name = strings.ToLower(scan.Step(
		"Name of the new hive micro-service based application",
		"required,lte=20,gte=3",
		scan.IsAlphanumDash))

	hive.FillMeta()

	return hive
}
