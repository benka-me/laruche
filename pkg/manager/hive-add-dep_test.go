package manager

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"testing"
)

func TestHiveAddDependencies(t *testing.T) {
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
			name: "dddd, eeee",
			args: args{
				hive:       &laruche.Hive{Name: "example", Author: "benka-me"},
				namespaces: laruche.Namespaces{"benka-me/dddd", "benka-me/eeee"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HiveAddDependencies(tt.args.hive, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("HiveAddDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
