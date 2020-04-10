package laruche

type Beez []*Bee
type BeezIter func(int, *Bee) error

func (beez *Beez) Map(fn BeezIter) error {
	for i, bee := range *beez {
		err := fn(i, bee)
		if err != nil {
			return err
		}
	}
	return nil
}

func (beez *Beez) Contain(needle Namespace) bool {
	flag := false
	beez.Map(func(i int, bee *Bee) error {
		if bee.GetNamespace() == needle {
			flag = true
		}
		return nil
	})
	return flag
}

func (beez *Beez) Find(needle Namespace) *Bee {
	var ret *Bee
	beez.Map(func(i int, bee *Bee) error {
		if bee.GetNamespace() == needle {
			ret = bee
			return nil
		}
		return nil
	})
	return ret
}

func (beez *Beez) Push(new *Bee) *Beez {
	if !beez.Contain(new.GetNamespace()) {
		*beez = append(*beez, new)
	}
	return beez
}

func (beez *Beez) GetDependencies() Namespaces {
	ret := make(Namespaces, len(*beez))
	beez.Map(func(i int, bee *Bee) error {
		ret[i] = bee.GetNamespace()
		return nil
	})
	return ret
}
