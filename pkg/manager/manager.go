package manager

import "github.com/benka-me/laruche/pkg/laruche"

type Context struct {
	Traversed laruche.Namespaces
	Consumers laruche.Namespaces
}

func newContext(oneOf laruche.OneOf) Context {
	ret := Context{
		Traversed: make(laruche.Namespaces, 0),
		Consumers: oneOf.GetConsumers(),
	}

	return ret
}
