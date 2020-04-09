package laruche

import (
	"reflect"
	"testing"
)

func TestNamespaces_Append(t *testing.T) {
	type args struct {
		src []Namespace
	}
	tests := []struct {
		name       string
		namespaces Namespaces
		args       args
		want       *Namespaces
	}{
		{
			name:       "regular 1",
			namespaces: Namespaces{"aaaa", "bbbb", "cccc"},
			args:       args{src: Namespaces{"dddd", "eeee"}},
			want:       &Namespaces{"aaaa", "bbbb", "cccc", "dddd", "eeee"},
		},
		{
			name:       "regular 2",
			namespaces: Namespaces{"aaaa", "bbbb", "cccc"},
			args:       args{src: Namespaces{}},
			want:       &Namespaces{"aaaa", "bbbb", "cccc"},
		},
		{
			name:       "empty receiver",
			namespaces: nil,
			args:       args{src: Namespaces{}},
			want:       &Namespaces{},
		},
		{
			name:       "empty receiver 2",
			namespaces: nil,
			args:       args{src: Namespaces{"nnnn"}},
			want:       &Namespaces{"nnnn"},
		},
		{
			name:       "empty both",
			namespaces: nil,
			args:       args{src: nil},
			want:       &Namespaces{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.namespaces.Append(tt.args.src...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}
