package laruche

type OneOf interface {
	GetConsumers() Namespaces
}

type OneOfArray interface {
	Map()
}

func (bee *Bee) GetConsumers() Namespaces {
	if bee == nil {
		return make(Namespaces, 0)
	}
	ret := make(Namespaces, len(bee.Cons))
	for i, c := range bee.Cons {
		ret[i] = Namespace(c)
	}
	return ret
}

func (hive *Hive) GetConsumers() Namespaces {
	if hive == nil {
		return make(Namespaces, 0)
	}
	ret := make(Namespaces, 1)
	ret[0] = hive.GetNamespace()
	return ret
}
