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
