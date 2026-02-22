package updater

import (
	"runtime"
	"testing"
)

func TestNew(t *testing.T) {
	u := New("1.0.0", false)
	if u == nil {
		t.Fatal("New() returned nil")
	}
	if u.currentVersion != "1.0.0" {
		t.Errorf("Expected version 1.0.0, got %s", u.currentVersion)
	}
	if u.quiet != false {
		t.Error("Expected quiet to be false")
	}
}

func TestGetBinaryName(t *testing.T) {
	u := New("1.0.0", true)

	tests := []struct {
		name     string
		expected string
	}{
		{"linux/amd64", "polyglot-linux-amd64"},
		{"linux/arm64", "polyglot-linux-arm64"},
		{"darwin/amd64", "polyglot-darwin-amd64"},
		{"darwin/arm64", "polyglot-darwin-arm64"},
		{"windows/amd64", "polyglot-windows-amd64.exe"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := u.getBinaryName()
			// We can only test the current platform
			if runtime.GOOS+"/"+runtime.GOARCH == tt.name {
				if result != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected, result)
				}
			}
		})
	}
}

func TestShouldCheckForUpdates(t *testing.T) {
	u := New("1.0.0", true)

	// First check should return true (no check file exists)
	if !u.ShouldCheckForUpdates() {
		t.Error("First check should return true")
	}
}

func TestIsInstalledViaHomebrew(t *testing.T) {
	u := New("1.0.0", true)

	// This will return false in test environment
	result := u.isInstalledViaHomebrew()
	if result {
		t.Log("Detected Homebrew installation")
	} else {
		t.Log("Not a Homebrew installation")
	}
}

func TestIsInstalledViaApt(t *testing.T) {
	u := New("1.0.0", true)

	// This will depend on the system
	result := u.isInstalledViaApt()
	if result {
		t.Log("Detected APT installation")
	} else {
		t.Log("Not an APT installation")
	}
}
