package cli

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/laruche/faker"
	"github.com/urfave/cli"
	"reflect"
	"testing"
)

func TestBee_AddDep(t *testing.T) {
	type args struct {
		depMode    bool
		namespaces laruche.Namespaces
	}
	tests := []struct {
		name    string
		bee     Bee
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bee.AddDep(tt.args.depMode, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("AddDep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOneOfInCurrentDir(t *testing.T) {
	tests := []struct {
		name    string
		want    OneOf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOneOfInCurrentDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneOfInCurrentDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOneOfInCurrentDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHive_AddDep(t *testing.T) {
	type args struct {
		depMode    bool
		namespaces laruche.Namespaces
	}
	tests := []struct {
		name    string
		hive    Hive
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.hive.AddDep(tt.args.depMode, tt.args.namespaces); (err != nil) != tt.wantErr {
				t.Errorf("AddDep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		app *App
	}
	tests := []struct {
		name string
		args args
		want cli.ActionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initBee(t *testing.T) {
	type test struct {
		args struct {
			app *App
		}
		name string
		want cli.ActionFunc
	}
	tests := make([]test, 0)
	testBees := faker.FakeBees(faker.AlphabetSeed)

	testBees.Map(func(i int, bee *laruche.Bee) error {
		tests = append(tests, test{
			name: bee.GetNamespaceStr(),
			want: nil,
		})
		return nil
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initBee(tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initBee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initHive(t *testing.T) {
	type args struct {
		app *App
	}
	tests := []struct {
		name string
		args args
		want cli.ActionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initHive(tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initHive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setAuthor(t *testing.T) {
	type args struct {
		hive *laruche.Hive
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
