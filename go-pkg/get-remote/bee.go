package remote

import (
	"context"
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/benka-me/users/go-pkg/users"
)

func GetBee(namespace laruche.Namespace) (*laruche.Bee, error) {
	res, err := clients.LarsrvGateway.GetBee(context.TODO(), &laruche.BeeReq{
		BeeName: namespace.String(),
		Token:   &users.Token{Val: config.GetState().AuthToken}})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func SaveBee(bee *laruche.Bee) error {
	return nil
}
