package laruche

func (b *Bee) GetSubDependencies() Namespaces {
	namespaces := make(Namespaces, len(b.Deps))
	for i, url := range b.Deps {
		namespaces[i] = Namespace(url)
	}
	return namespaces
}

func (h *Hive) GetDependencies() Namespaces {
	namespaces := make(Namespaces, len(h.Deps))
	i := 0
	for url := range h.Deps {
		namespaces[i] = Namespace(url)
	}
	return namespaces
}
