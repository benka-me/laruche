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
