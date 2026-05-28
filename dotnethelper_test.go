package dotnethelper

import (
	"os/exec"
	"testing"
)

func skipIfNoDotnet(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("dotnet"); err != nil {
		t.Skipf("dotnet not found in PATH: %v", err)
	}
}

func TestInstalled(t *testing.T) {
	skipIfNoDotnet(t)
	tests := []struct {
		name string
		want bool
	}{
		{name: "1",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Installed(); got != tt.want {
				t.Errorf("Installed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion(t *testing.T) {
	skipIfNoDotnet(t)
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Version()
			if (err != nil) != tt.wantErr {
				t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("empty Version()")
			}
		})
	}
}
