package config

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetHive(namespace string) (*laruche.Hive, error) {
	hive := laruche.Hive{}
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/hive.yaml", LaruchePath, namespace))
	if err != nil {
		return &hive, err
	}

	err = yaml.Unmarshal(dat, &hive)
	return &hive, err
}
