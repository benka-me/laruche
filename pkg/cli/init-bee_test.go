package cli

import (
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"github.com/benka-me/laruche/pkg/local"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func reset() {
	config.RemoveAllBee()
}

func Test_InitAlphaBees(t *testing.T) {
	reset()

	type test struct {
		arg     *laruche.Bee
		name    string
		wantErr bool
	}
	tests := make([]test, 0)
	faker.Alpha.Map(func(i int, bee *laruche.Bee) error {
		tests = append(tests, test{
			name:    bee.GetNamespaceStr(),
			arg:     bee,
			wantErr: false,
		})
		return nil
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := actionInitBee(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("actionInitBee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		t.Run(tt.name+"/cmp-local", func(t *testing.T) {
			if ret, err := local.GetBee(tt.arg.GetNamespaceStr()); !cmp.Equal(ret, tt.arg) || err != nil {
				t.Errorf("cmp.Equal() ret = %v, tt.arg %v, err %v", ret, tt.arg, err)
			}
		})
	}
}
