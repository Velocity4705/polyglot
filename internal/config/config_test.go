package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}

	if cfg.AutoInstall.Enabled {
		t.Error("Default auto_install.enabled should be false")
	}

	if cfg.AutoInstall.Method != "system" {
		t.Errorf("Default auto_install.method should be 'system', got '%s'", cfg.AutoInstall.Method)
	}

	if !cfg.AutoInstall.Prompt {
		t.Error("Default auto_install.prompt should be true")
	}

	if cfg.Languages == nil {
		t.Error("Languages map should be initialized")
	}

	if cfg.CustomExtensions == nil {
		t.Error("CustomExtensions map should be initialized")
	}

	if cfg.Environment == nil {
		t.Error("Environment map should be initialized")
	}
}

func TestLanguageConfig(t *testing.T) {
	cfg := DefaultConfig()

	// Test setting language config
	cfg.Languages["python"] = LanguageConfig{
		Enabled: true,
		Command: "python3.11",
		Version: "3.11",
	}

	langCfg := cfg.GetLanguageConfig("python")
	if langCfg == nil {
		t.Fatal("GetLanguageConfig returned nil")
	}

	if !langCfg.Enabled {
		t.Error("Python should be enabled")
	}

	if langCfg.Command != "python3.11" {
		t.Errorf("Expected command 'python3.11', got '%s'", langCfg.Command)
	}

	if langCfg.Version != "3.11" {
		t.Errorf("Expected version '3.11', got '%s'", langCfg.Version)
	}
}

func TestCustomExtensions(t *testing.T) {
	cfg := DefaultConfig()

	// Test setting custom extension
	cfg.CustomExtensions[".pyx"] = "Python"

	ext := cfg.GetCustomExtension(".pyx")
	if ext != "Python" {
		t.Errorf("Expected extension mapping to 'Python', got '%s'", ext)
	}

	// Test non-existent extension
	ext = cfg.GetCustomExtension(".xyz")
	if ext != "" {
		t.Errorf("Expected empty string for non-existent extension, got '%s'", ext)
	}
}

func TestEnvironmentVariables(t *testing.T) {
	cfg := DefaultConfig()

	// Set global environment
	cfg.Environment["GLOBAL_VAR"] = "global_value"

	// Set language-specific environment
	cfg.Languages["python"] = LanguageConfig{
		Enabled: true,
		Environment: map[string]string{
			"PYTHON_VAR": "python_value",
			"GLOBAL_VAR": "python_override",
		},
	}

	// Test global environment
	env := cfg.GetEnvironment("nonexistent")
	if env["GLOBAL_VAR"] != "global_value" {
		t.Errorf("Expected GLOBAL_VAR='global_value', got '%s'", env["GLOBAL_VAR"])
	}

	// Test language-specific environment (should override global)
	env = cfg.GetEnvironment("python")
	if env["PYTHON_VAR"] != "python_value" {
		t.Errorf("Expected PYTHON_VAR='python_value', got '%s'", env["PYTHON_VAR"])
	}
	if env["GLOBAL_VAR"] != "python_override" {
		t.Errorf("Expected GLOBAL_VAR='python_override', got '%s'", env["GLOBAL_VAR"])
	}
}

func TestIsLanguageEnabled(t *testing.T) {
	cfg := DefaultConfig()

	// Test default (should be enabled)
	if !cfg.IsLanguageEnabled("python") {
		t.Error("Language should be enabled by default")
	}

	// Test explicitly disabled
	cfg.Languages["python"] = LanguageConfig{
		Enabled: false,
	}
	if cfg.IsLanguageEnabled("python") {
		t.Error("Language should be disabled")
	}

	// Test explicitly enabled
	cfg.Languages["python"] = LanguageConfig{
		Enabled: true,
	}
	if !cfg.IsLanguageEnabled("python") {
		t.Error("Language should be enabled")
	}
}

func TestConfigMerging(t *testing.T) {
	base := &Config{
		AutoInstall: AutoInstallConfig{
			Enabled: false,
			Method:  "system",
			Prompt:  true,
		},
		Languages: map[string]LanguageConfig{
			"python": {
				Enabled: true,
				Version: "3.9",
			},
		},
		CustomExtensions: map[string]string{
			".pyx": "Python",
		},
		Environment: map[string]string{
			"VAR1": "base_value",
		},
	}

	override := &Config{
		AutoInstall: AutoInstallConfig{
			Enabled: true,
			Method:  "script",
			Prompt:  false,
		},
		Languages: map[string]LanguageConfig{
			"python": {
				Enabled: true,
				Version: "3.11",
			},
			"javascript": {
				Enabled: true,
			},
		},
		CustomExtensions: map[string]string{
			".jsx": "JavaScript",
		},
		Environment: map[string]string{
			"VAR2": "override_value",
		},
	}

	result := mergeConfigs(base, override)

	// Test auto_install override
	if !result.AutoInstall.Enabled {
		t.Error("AutoInstall.Enabled should be overridden to true")
	}
	if result.AutoInstall.Method != "script" {
		t.Errorf("AutoInstall.Method should be 'script', got '%s'", result.AutoInstall.Method)
	}

	// Test language merging
	if result.Languages["python"].Version != "3.11" {
		t.Errorf("Python version should be '3.11', got '%s'", result.Languages["python"].Version)
	}
	if !result.Languages["javascript"].Enabled {
		t.Error("JavaScript should be enabled")
	}

	// Test custom extensions merging
	if result.CustomExtensions[".pyx"] != "Python" {
		t.Error("Base custom extension should be preserved")
	}
	if result.CustomExtensions[".jsx"] != "JavaScript" {
		t.Error("Override custom extension should be added")
	}

	// Test environment merging
	if result.Environment["VAR1"] != "base_value" {
		t.Error("Base environment variable should be preserved")
	}
	if result.Environment["VAR2"] != "override_value" {
		t.Error("Override environment variable should be added")
	}
}

func TestSaveAndLoad(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "polyglot-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Override config path for testing
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// Create config
	cfg := DefaultConfig()
	cfg.AutoInstall.Enabled = true
	cfg.Languages["python"] = LanguageConfig{
		Enabled: true,
		Version: "3.11",
	}

	// Save config
	err = cfg.Save()
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Check file exists
	configPath := filepath.Join(tmpDir, ".polyglot", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("Config file was not created")
	}

	// Load config
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify loaded config
	if !loaded.AutoInstall.Enabled {
		t.Error("Loaded config should have auto_install enabled")
	}

	langCfg := loaded.GetLanguageConfig("python")
	if langCfg == nil || langCfg.Version != "3.11" {
		t.Error("Loaded config should have Python version 3.11")
	}
}
