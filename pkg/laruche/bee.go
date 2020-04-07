package laruche

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

func InitBeeAskUser() *Bee {
	bee := &Bee{}
	scan.V = validator.New()

	bee.Name = strings.ToLower(scan.Step(
		"Name of the new bee microservice ",
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

func (b *Bee) GetLocal() *Bee {
	if b == nil {
		return b
	}
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/src/%s/bee.yaml", "", b.Repo))
	if err != nil {
		return b
	}

	_ = yaml.Unmarshal(dat, b)
	return b
}

func GetLocalBeeCurrentDir() (Bee, error) {
	bee := Bee{}
	dat, err := ioutil.ReadFile("./bee.yaml")

	if err != nil {
		return bee, err
	}

	err = yaml.Unmarshal(dat, &bee)
	if err != nil {
		return bee, err
	}

	return bee, nil
}

func (bee *Bee) SaveLocal() error {
	data, err := yaml.Marshal(bee)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/src/%s/bee.yaml", "", bee.Repo)
	err = ioutil.WriteFile(path, data, 0755)
	if err != nil {
		return err
	}

	conf.AddBee(bee.GetNamespaceStr(), bee.Repo)
	return nil
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
