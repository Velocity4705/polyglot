package installer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Installer handles toolchain installation
type Installer struct {
	pkgManager *PackageManager
	autoYes    bool
	quiet      bool
}

// New creates a new Installer
func New(autoYes, quiet bool) *Installer {
	return &Installer{
		pkgManager: DetectPackageManager(),
		autoYes:    autoYes,
		quiet:      quiet,
	}
}

// IsAvailable checks if a toolchain is available
func (i *Installer) IsAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// PromptInstall prompts the user to install a toolchain
func (i *Installer) PromptInstall(language, command string) bool {
	if i.autoYes {
		return true
	}

	if i.pkgManager == nil {
		fmt.Printf("\n⚠️  %s not found in your system\n", language)
		fmt.Println("\nNo package manager detected.")
		fmt.Printf("Please install %s manually:\n", language)
		i.showManualInstructions(command)
		return false
	}

	pkgName := GetPackageName(command, i.pkgManager.Name)

	fmt.Printf("\n⚠️  %s not found in your system\n", language)
	fmt.Printf("\nPolyglot can install %s for you:\n", language)
	fmt.Printf("  • Installation method: System package manager (%s)\n", i.pkgManager.Name)
	fmt.Printf("  • Package: %s\n", pkgName)
	fmt.Printf("  • Will be installed system-wide (available to all programs)\n")

	if needsSudo(i.pkgManager.Name) {
		fmt.Printf("  • Requires: sudo password\n")
	}

	fmt.Printf("\nInstall %s now? [Y/n] ", language)

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	return response == "" || response == "y" || response == "yes"
}

// Install installs a toolchain
func (i *Installer) Install(language, command string) error {
	if i.pkgManager == nil {
		return fmt.Errorf("no package manager detected")
	}

	pkgName := GetPackageName(command, i.pkgManager.Name)

	if !i.quiet {
		fmt.Printf("\n📦 Installing %s...\n", language)
		fmt.Printf("   Running: %s %s\n", i.pkgManager.Name, pkgName)
	}

	// Build install command
	installCmd := append(i.pkgManager.InstallCmd, pkgName)

	// Execute installation
	cmd := exec.Command(installCmd[0], installCmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}

	if !i.quiet {
		fmt.Printf("\n✅ %s installed successfully!\n", language)
		fmt.Printf("💡 %s is now available system-wide\n", language)
	}

	return nil
}

// InstallIfMissing checks if a toolchain is available and installs if missing
func (i *Installer) InstallIfMissing(language, command string) error {
	// Check if already installed
	if i.IsAvailable(command) {
		return nil
	}

	// Prompt user
	if !i.PromptInstall(language, command) {
		return fmt.Errorf("%s not installed", language)
	}

	// Install
	return i.Install(language, command)
}

// showManualInstructions shows manual installation instructions
func (i *Installer) showManualInstructions(command string) {
	fmt.Println("\nManual installation instructions:")

	switch GetOS() {
	case "linux":
		fmt.Println("\nFedora/RHEL:")
		fmt.Printf("  sudo dnf install %s\n", GetPackageName(command, "dnf"))
		fmt.Println("\nUbuntu/Debian:")
		fmt.Printf("  sudo apt install %s\n", GetPackageName(command, "apt"))
		fmt.Println("\nArch:")
		fmt.Printf("  sudo pacman -S %s\n", GetPackageName(command, "pacman"))
	case "darwin":
		fmt.Println("\nmacOS:")
		fmt.Printf("  brew install %s\n", GetPackageName(command, "brew"))
	default:
		fmt.Printf("\nPlease install %s for your system\n", command)
	}

	fmt.Println("\nAfter installation, run: polyglot check")
}

// needsSudo checks if a package manager needs sudo
func needsSudo(pkgManager string) bool {
	return pkgManager != "brew"
}
