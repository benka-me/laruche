package discover

import (
	"errors"
	"fmt"
	"github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"google.golang.org/grpc"
	"strconv"
)

type Engine struct {
	Dev         bool
	Deps        map[string]*laruche.Dep
	GatewayPort int
}

func (r *Engine) ThisPort(namespace string) (string, error) {
	if _, ok := r.Deps[namespace]; !ok {
		fmt.Println()
		return "", errors.New("laruche engine: can't find this port because" + namespace + " is not registered")
	}
	i := r.Deps[namespace].Port
	return ":" + strconv.FormatInt(int64(i), 10), nil
}

func (r *Engine) GrpcConn(namespace string, gateway bool, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	var host string
	if r.Dev {
		host = r.Deps[namespace].Dev
	} else {
		host = r.Deps[namespace].Prod
	}

	var port string
	var isGateway string
	if gateway {
		port = strconv.FormatInt(int64(r.GatewayPort), 10)
		isGateway = "(Through gateway)"
	} else {
		port = strconv.FormatInt(int64(r.Deps[namespace].Port), 10)
	}
	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("New client to %s service on %s %s\n", namespace, address, isGateway)
	return grpc.Dial(address, options...)
}

func ParseEngine(namespace string, dev bool) (*Engine, error) {
	hive, err := local.GetHive(laruche.Namespace(namespace))
	if err != nil {
		return &Engine{}, err
	}

	return &Engine{
		Dev:  dev,
		Deps: hive.Deps,
	}, nil
}
