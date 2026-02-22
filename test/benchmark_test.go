package test

import (
	"os/exec"
	"testing"
)

// BenchmarkDetection benchmarks language detection
func BenchmarkDetection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(polyglotBinary, "run", "--dry-run", "fixtures/hello.py")
		_ = cmd.Run()
	}
}

// BenchmarkPythonExecution benchmarks Python execution
func BenchmarkPythonExecution(b *testing.B) {
	// Check if python3 is available
	if _, err := exec.LookPath("python3"); err != nil {
		b.Skip("python3 not available")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(polyglotBinary, "run", "-q", "fixtures/hello.py")
		_ = cmd.Run()
	}
}

// BenchmarkGoExecution benchmarks Go execution
func BenchmarkGoExecution(b *testing.B) {
	// Check if go is available
	if _, err := exec.LookPath("go"); err != nil {
		b.Skip("go not available")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(polyglotBinary, "run", "-q", "fixtures/hello.go")
		_ = cmd.Run()
	}
}

// BenchmarkJavaScriptExecution benchmarks JavaScript execution
func BenchmarkJavaScriptExecution(b *testing.B) {
	// Check if node is available
	if _, err := exec.LookPath("node"); err != nil {
		b.Skip("node not available")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(polyglotBinary, "run", "-q", "fixtures/hello.js")
		_ = cmd.Run()
	}
}

// BenchmarkConfigLoad benchmarks configuration loading
func BenchmarkConfigLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(polyglotBinary, "config", "show")
		_ = cmd.Run()
	}
}

// BenchmarkListCommand benchmarks the list command
func BenchmarkListCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(polyglotBinary, "list")
		_ = cmd.Run()
	}
}
