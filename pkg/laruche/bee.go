package laruche

type Beez []*Bee
type BeezIter func(int, *Bee) error

func (beez *Beez) Map(fn BeezIter) {
	for i, bee := range *beez {
		err := fn(i, bee)
		if err != nil {
			return
		}
	}
}
