package manager_test

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"github.com/benka-me/laruche/pkg/manager"
	"testing"
)

func TestBeeAddDependencies(t *testing.T) {
	//config.RemoveAllBee()
	//faker.AlphaBees.Map(func(i int, bee *laruche.Bee) error {
	//	_ = cli.ActionInitBee(bee)
	//	return nil
	//})
	type args struct {
		bee        *laruche.Bee
		namespaces laruche.Namespaces
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "d -> e",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/dddd"],
				namespaces: laruche.Namespaces{"benka-me/eeee"},
			},
			wantErr: false,
		},
	}
	tests[0].args.bee.Deps = []string{"benka-me/eeee"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := manager.BeeAddDependencies(tt.args.bee, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("BeeAddDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
