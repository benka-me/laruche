package generator

import "testing"

func TestPkgNames_Import(t *testing.T) {
	tests := []struct {
		name         string
		PkgName      string
		PkgNameCamel string
		want         string
	}{
		{
			name: "One",
			want: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := PkgNames{
				PkgName:      tt.PkgName,
				PkgNameCamel: tt.PkgNameCamel,
			}
			if got := cl.Import(); got != tt.want {
				t.Errorf("Import() = %v, want %v", got, tt.want)
			}
		})
	}
}
