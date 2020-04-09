package manager

import "github.com/benka-me/laruche/pkg/laruche"

type Context struct {
	Traversed laruche.Namespaces
	Consumers laruche.Namespaces
}
