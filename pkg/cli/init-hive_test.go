package cli

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"testing"
)

func TestActionInitHive(t *testing.T) {
	type args struct {
		hive *laruche.Hive
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "example",
			args: args{hive: &laruche.Hive{
				Name:   "example",
				Author: "benka-me",
				Public: false,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ActionInitHive(tt.args.hive); (err != nil) != tt.wantErr {
				t.Errorf("ActionInitHive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
