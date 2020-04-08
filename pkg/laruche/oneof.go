package laruche

type OneOf interface {
	AddDep(bool, Namespaces) error
}

func (bee *Bee) AddDep(depMode bool, namespaces Namespaces) error {
	if depMode {
		namespaces = bee.GetSubDependencies()
	}

	return nil
}

func (hive *Hive) AddDep(depMode bool, namespaces Namespaces) error {
	if depMode {
		namespaces = hive.GetDependencies()
	}
	return nil
}
