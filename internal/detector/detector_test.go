package detector

import (
	"testing"

	"github.com/Velocity4705/polyglot/pkg/types"
)

func TestDetectLanguage(t *testing.T) {
	tests := []struct {
		filename string
		expected types.Language
		wantErr  bool
	}{
		{"hello.py", types.Python, false},
		{"main.go", types.Go, false},
		{"app.js", types.JavaScript, false},
		{"script.rb", types.Ruby, false},
		{"index.php", types.PHP, false},
		{"Hello.java", types.Java, false},
		{"program.c", types.C, false},
		{"program.cpp", types.Cpp, false},
		{"program.cc", types.Cpp, false},
		{"main.rs", types.Rust, false},
		{"unknown.xyz", types.Unknown, true},
		{"noextension", types.Unknown, true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := DetectLanguage(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectLanguage(%q) error = %v, wantErr %v", tt.filename, err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("DetectLanguage(%q) = %v, want %v", tt.filename, got, tt.expected)
			}
		})
	}
}

func TestDetectLanguageCaseInsensitive(t *testing.T) {
	tests := []struct {
		filename string
		expected types.Language
	}{
		{"HELLO.PY", types.Python},
		{"Main.GO", types.Go},
		{"App.JS", types.JavaScript},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := DetectLanguage(tt.filename)
			if err != nil {
				t.Errorf("DetectLanguage(%q) unexpected error: %v", tt.filename, err)
				return
			}
			if got != tt.expected {
				t.Errorf("DetectLanguage(%q) = %v, want %v", tt.filename, got, tt.expected)
			}
		})
	}
}
