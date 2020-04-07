package generator

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"testing"
)

func TestGo_GenerateServer(t *testing.T) {
	tests := []struct {
		name    string
		Setup   *laruche.LanguageSetup
		bee     *laruche.Bee
		wantErr bool
	}{
		{
			name: "sub0",
			Setup: &laruche.LanguageSetup{
				Active: true,
				//Repo:    "github.com/benka-me/sub0",
				//Files:   []string{
				//	"sub0.proto",
				//	"rpc-sub0.proto",
				//},
				//PkgName: "sub0",
			},
			bee: &laruche.Bee{
				DevLang: laruche.DevLang(0),
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Go{
				Setup: tt.Setup,
			}
			if err := g.ClientsFile(tt.bee); (err != nil) != tt.wantErr {
				t.Errorf("ClientsFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
