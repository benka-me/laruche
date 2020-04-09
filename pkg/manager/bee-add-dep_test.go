package manager

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"testing"
)

func TestBeeAddDependencies(t *testing.T) {
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
			name: "regular 1",
			args: args{
				bee: &laruche.Bee{
					Name:         "",
					PkgName:      "",
					PkgNameCamel: "",
					Repo:         "",
					Author:       "",
					Port:         0,
					Public:       false,
					License:      "",
					Description:  "",
					Keywords:     "",
					Tag:          "",
					DevLang:      0,
					Languages:    nil,
					ProtoSetup:   nil,
					IsGateway:    false,
					Deps:         nil,
					Cons:         nil,
				},
				namespaces: nil,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BeeAddDependencies(tt.args.bee, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("BeeAddDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
