package generator

import (
	"github.com/benka-me/laruche/pkg/laruche"
)

type LangGenerator interface {
	Protoc(*laruche.Bee)
	ClientsFile(*laruche.Bee) error
	ServerFiles(*laruche.Bee) error
}

type LangGenerators *[]LangGenerator

func GetLangs(lgs *laruche.Languages) (LangGenerators, error) {
	var langs = make([]LangGenerator, 2)
	langs[0] = Go(*lgs.GetGo())
	langs[1] = Javascript(*lgs.GetJavascript())

	return &langs, nil
}

func GetDevLang(b *laruche.Bee) LangGenerator {
	var EnumLang = map[laruche.DevLang]LangGenerator{
		0: Go(*b.Languages.Go),
		1: Javascript(*b.Languages.Javascript),
	}
	return EnumLang[b.DevLang]
}
