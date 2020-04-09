package cli

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"testing"
)

func Test_actionInitBee(t *testing.T) {

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
			wantErr: nil,
		})
		return nil
	})
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if err := actionInitBee(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("actionInitBee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
