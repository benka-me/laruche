package laruche

type MBees map[string]*Bee
type MBeeFunc func(string, *Bee) interface{}

func (m MBees) Map(fn MBeeFunc) MBees {
	if m == nil {
		return make(MBees)
	}
	for namespace, bee := range m {
		fn(namespace, bee)
	}
	return m
}

type NamespaceIter func(int, Namespace) error

func (namespaces Namespaces) Map(fn NamespaceIter) error {
	for i, n := range namespaces {
		err := fn(i, n)
		if err != nil {
			return err
		}
	}
	return nil
}
