package test

import (
	"github.com/benka-me/laruche/pkg/cli"
	"github.com/benka-me/laruche/pkg/config"
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func initHiveExampleA(t *testing.T) {
	type args struct {
		hive *laruche.Hive
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "example-a",
			args: args{hive: &laruche.Hive{
				Name:   "example-a",
				Author: "benka-me",
				Public: false,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cli.ActionInitHive(tt.args.hive); (err != nil) != tt.wantErr {
				t.Errorf("ActionInitHive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func initHiveExampleB(t *testing.T) {
	type args struct {
		hive *laruche.Hive
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "example-b",
			args: args{hive: &laruche.Hive{
				Name:   "example-b",
				Author: "benka-me",
				Public: false,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cli.ActionInitHive(tt.args.hive); (err != nil) != tt.wantErr {
				t.Errorf("ActionInitHive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func initAlphaBees(t *testing.T) {
	config.RemoveAllBee()

	type test struct {
		arg     *laruche.Bee
		name    string
		wantErr bool
	}
	tests := make([]test, 0)
	faker.AlphaBees.Map(func(i int, bee *laruche.Bee) error {
		tests = append(tests, test{
			name:    bee.GetNamespaceStr(),
			arg:     bee,
			wantErr: false,
		})
		return nil
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cli.ActionInitBee(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("actionInitBee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		t.Run(tt.name+"/cmp-local", func(t *testing.T) {
			if ret, err := local.GetBee(tt.arg.GetNamespace()); !cmp.Equal(ret, tt.arg) || err != nil {
				t.Errorf("cmp.Equal() ret = %v, tt.arg %v, err %v", ret, tt.arg, err)
			}
		})
	}
}
