package remote

import (
	"github.com/benka-me/laruche/go-pkg/http/rpc"
	"google.golang.org/grpc"
)

var clients = rpc.InitClients(grpc.WithInsecure())
