package laruche

import (
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"strings"
)

type Hives []*Hive

func (hive *Hive) FillMeta() {
	hive.PkgNameCamel = strings.Title(scan.KebabToCamelCase(hive.PkgName))
}
