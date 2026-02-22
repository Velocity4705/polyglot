class Polyglot < Formula
  desc "Universal compiler and interpreter wrapper for 30+ programming languages"
  homepage "https://github.com/yourusername/polyglot"
  url "https://github.com/yourusername/polyglot/archive/v0.6.0.tar.gz"
  sha256 "REPLACE_WITH_ACTUAL_SHA256"
  license "MIT"
  head "https://github.com/yourusername/polyglot.git", branch: "main"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/polyglot"
    
    # Install shell completions
    generate_completions_from_executable(bin/"polyglot", "completion")
  end

  test do
    assert_match "Polyglot v#{version}", shell_output("#{bin}/polyglot version")
    
    # Test with a simple Python script
    (testpath/"hello.py").write "print('Hello from Polyglot!')"
    assert_match "Hello from Polyglot!", shell_output("#{bin}/polyglot run --dry-run #{testpath}/hello.py")
  end
end
