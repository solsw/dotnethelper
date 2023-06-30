package dotnethelper

import (
	"os/exec"
	"strings"
)

// Installed reports whether .NET is installed.
func Installed() bool {
	return exec.Command("dotnet", "--version").Run() == nil
}

// Version returns version (as returned by 'dotnet --version') of installed .NET, if any.
func Version() (string, error) {
	o, err := exec.Command("dotnet", "--version").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(o)), nil
}
