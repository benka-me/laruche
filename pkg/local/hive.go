package local

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/laruche"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func GetHive(namespace string) (*laruche.Hive, error) {
	hive := laruche.Hive{}
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/hive.yaml", config.LaruchePath, namespace))
	if err != nil {
		return &hive, err
	}

	err = yaml.Unmarshal(dat, &hive)
	return &hive, err
}

func GetHiveCurrentDir() (*laruche.Hive, error) {
	hive := &laruche.Hive{}
	dat, err := ioutil.ReadFile("./hive.yaml")

	if err != nil {
		return hive, err
	}

	err = yaml.Unmarshal(dat, &hive)
	if err != nil {
		return hive, err
	}

	return hive, nil
}

func SaveLocal(hive *laruche.Hive) error {
	data, err := yaml.Marshal(hive)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/%s", config.LaruchePath, hive.GetNamespaceStr()), 0755) //TODO perm
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/hive.yaml", config.LaruchePath, hive.GetNamespaceStr())
	err = ioutil.WriteFile(path, data, 0755)
	if err != nil {
		return err
	}

	return nil
}
