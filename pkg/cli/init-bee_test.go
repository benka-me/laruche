package cli

import (
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_InitAlphaBees(t *testing.T) {
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
			if err := ActionInitBee(tt.arg); (err != nil) != tt.wantErr {
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
