package faker

import (
	"github.com/benka-me/laruche/pkg/cli/form"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/brianvoe/gofakeit"
)

var Alpha = FakeBees(AlphabetSeed)
var Sample1 = FakeBees(Sample1Seed)

func FakeBees(bees []*laruche.Bee) laruche.Beez {
	ret := make([]*laruche.Bee, len(bees))

	for i, bee := range bees {
		form.FillDefaultMeta(bee)
		bee.Repo = "test/" + bee.Name
		bee.Author = "benka-me"
		bee.Port = int32(gofakeit.Number(200, 65000))
		bee.Public = true
		bee.License = gofakeit.RandString([]string{"MIT", "AGPL2", "AGPL3", "Apache"})
		bee.Description = gofakeit.HipsterSentence(23)
		bee.Keywords = gofakeit.HipsterSentence(10)
		bee.DevLang = laruche.DevLang(gofakeit.Number(0, 0))
		bee.IsGateway = false
		bee.Deps = []string{}
		bee.Cons = []string{}
		ret[i] = bee
	}
	return ret
}
func FakeMBees(bees []*laruche.Bee) laruche.MBees {
	ret := make(laruche.MBees)

	for _, bee := range bees {
		form.FillDefaultMeta(bee)
		bee.Repo = "test/" + bee.Name
		bee.Author = "benka-me"
		bee.Port = int32(gofakeit.Number(200, 65000))
		bee.Public = true
		bee.License = gofakeit.RandString([]string{"MIT", "AGPL2", "AGPL3", "Apache"})
		bee.Description = gofakeit.HipsterSentence(23)
		bee.Keywords = gofakeit.HipsterSentence(10)
		bee.DevLang = laruche.DevLang(gofakeit.Number(0, 0))
		bee.IsGateway = false
		bee.Deps = []string{}
		bee.Cons = []string{}
		ret[bee.GetNamespaceStr()] = bee
	}
	return ret
}
