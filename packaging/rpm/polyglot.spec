Name:           polyglot
Version:        0.6.0
Release:        1%{?dist}
Summary:        Universal compiler and interpreter wrapper

License:        MIT
URL:            https://github.com/yourusername/polyglot
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  golang >= 1.21
Requires:       glibc

%description
Polyglot is a universal compiler and interpreter wrapper that detects
programming languages by file extension and executes them with the
appropriate toolchain.

Features:
- Supports 30+ programming languages
- Automatic language detection
- Auto-install missing toolchains
- Watch mode for development
- Colored output
- Comprehensive configuration

%prep
%setup -q

%build
go build -ldflags="-s -w" -o %{name} ./cmd/polyglot

%install
install -D -m 0755 %{name} %{buildroot}%{_bindir}/%{name}

%files
%license LICENSE
%doc README.md
%{_bindir}/%{name}

%changelog
* Sat Feb 22 2026 Your Name <your.email@example.com> - 0.6.0-1
- Release 0.6.0
- Added comprehensive testing
- Added distribution packaging
- 80% milestone completion
