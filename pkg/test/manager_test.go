package test

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"github.com/benka-me/laruche/pkg/manager"
	"testing"
)

func Test1Init(t *testing.T) {
	initHiveExampleA(t)
	initAlphaBees(t)
}

func Test2BeeAddDependenciesHiveA(t *testing.T) {
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
			name: "g -> a",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/gggg"],
				namespaces: laruche.Namespaces{"benka-me/aaaa"},
			},
			wantErr: false,
		},
		{
			name: "d -> e/f",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/dddd"],
				namespaces: laruche.Namespaces{"benka-me/eeee", "benka-me/ffff"},
			},
			wantErr: false,
		},
		{
			name: "a -> d",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/aaaa"],
				namespaces: laruche.Namespaces{"benka-me/dddd"},
			},
			wantErr: false,
		},
		{
			name: "i -> j",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/iiii"],
				namespaces: laruche.Namespaces{"benka-me/jjjj"},
			},
			wantErr: false,
		},
		{
			name: "j -> g",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/jjjj"],
				namespaces: laruche.Namespaces{"benka-me/gggg"},
			},
			wantErr: false,
		},
		{
			name: "b -> f/g",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/bbbb"],
				namespaces: laruche.Namespaces{"benka-me/ffff", "benka-me/gggg"},
			},
			wantErr: false,
		},
		{
			name: "c -> h/i",
			args: args{
				bee:        faker.AlphaBeesMap["benka-me/cccc"],
				namespaces: laruche.Namespaces{"benka-me/iiii", "benka-me/hhhh"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := manager.BeeAddDependencies(tt.args.bee, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("BeeAddDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test3HiveAddDependencies(t *testing.T) {
	type args struct {
		hive       *laruche.Hive
		namespaces laruche.Namespaces
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "example-a -> bad namespace",
			args: args{
				hive:       &laruche.Hive{Name: "example-a", Author: "benka-me"},
				namespaces: laruche.Namespaces{"benka-me/ccccf"},
			},
			wantErr: true,
		},
		{
			name: "example-a -> c",
			args: args{
				hive:       &laruche.Hive{Name: "example-a", Author: "benka-me"},
				namespaces: laruche.Namespaces{"benka-me/cccc"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := manager.HiveAddDependencies(tt.args.hive, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("HiveAddDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
