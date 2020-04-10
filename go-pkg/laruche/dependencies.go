package laruche

func (bee Bee) GetDependencies() Namespaces {
	namespaces := make(Namespaces, len(bee.Deps))
	for i, url := range bee.Deps {
		namespaces[i] = Namespace(url)
	}
	return namespaces
}

func (hive Hive) GetDependencies() Namespaces {
	namespaces := make(Namespaces, 0)
	for url := range hive.Deps {
		namespaces = append(namespaces, Namespace(url))
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
func (bee *Bee) PushDependencies(namespaces Namespaces) *Bee {
	if bee.Deps == nil {
		bee.Deps = make([]string, 0)
	}
	bee.Deps = AppendUniqueString(bee.Deps, namespaces.String()...)
	return bee
}

func (bee *Bee) PushConsumer(namespace Namespace) *Bee {
	if bee.Cons == nil {
		bee.Cons = make([]string, 0)
	}
	bee.Cons = AppendUniqueString(bee.Cons, namespace.String())
	return bee
}
func (bee *Bee) PushConsumers(namespaces Namespaces) *Bee {
	if bee.Cons == nil {
		bee.Cons = make([]string, 0)
	}
	bee.Cons = AppendUniqueString(bee.Cons, namespaces.String()...)
	return bee
}
