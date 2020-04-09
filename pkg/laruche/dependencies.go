package laruche

func (bee Bee) GetSubDependencies() Namespaces {
	namespaces := make(Namespaces, len(bee.Deps))
	for i, url := range bee.Deps {
		namespaces[i] = Namespace(url)
	}
	return namespaces
}

func (hive Hive) GetDependencies() Namespaces {
	namespaces := make(Namespaces, len(hive.Deps))
	i := 0
	for url := range hive.Deps {
		namespaces[i] = Namespace(url)
	}
	return namespaces
}

func (bee *Bee) PushDependency(namespace Namespace) *Bee {
	if bee.Deps == nil {
		bee.Deps = make([]string, 0)
	}
	bee.Deps = append(bee.Deps, string(namespace))
	return bee
}
