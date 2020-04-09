package local

import (
	"errors"
	"fmt"
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/laruche"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetBee(namespace laruche.Namespace) (*laruche.Bee, error) {
	bee := laruche.Bee{}
	beeConf, err := config.GetBee(namespace)
	if err != nil {
		return &laruche.Bee{}, errors.New("bee doesn't exist on this machine")
	}
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/bee.yaml", config.SourcePath, beeConf.Path))
	if err != nil {
		return &bee, err
	}

	err = yaml.Unmarshal(dat, &bee)
	return &bee, err
}

func GetBeeCurrentDir() (*laruche.Bee, error) {
	bee := &laruche.Bee{}
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

func SaveBee(bee *laruche.Bee) error {
	data, err := yaml.Marshal(bee)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/bee.yaml", config.SourcePath, bee.Repo)
	err = ioutil.WriteFile(path, data, 0755)
	if err != nil {
		return err
	}
	return nil
}
