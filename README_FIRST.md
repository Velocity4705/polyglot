# 👋 Welcome to Polyglot!

**Universal Compiler Wrapper - One Command to Run Them All**

---

## 🚀 Quick Start (30 seconds)

```bash
# Build
make build

# Try it
./bin/polyglot run test/fixtures/hello.py
./bin/polyglot run test/fixtures/hello.go
./bin/polyglot run test/fixtures/hello.js

# Install
make install
```

---

## 📖 What is Polyglot?

Polyglot automatically detects programming languages and runs them with the right toolchain.

**Instead of:**
```bash
python3 script.py
go run main.go
node app.js
ruby script.rb
php index.php
```

**Just use:**
```bash
polyglot run script.py
polyglot run main.go
polyglot run app.js
polyglot run script.rb
polyglot run index.php
```

---

## ✨ Features

- ✅ **5 Languages**: Python, Go, JavaScript, Ruby, PHP
- ✅ **Auto-Detection**: Smart file extension recognition
- ✅ **Simple CLI**: Intuitive commands
- ✅ **Argument Passing**: `--args` flag
- ✅ **Quiet/Verbose**: `-q` and `-v` modes
- ✅ **Toolchain Check**: Verify installations
- ✅ **Production Ready**: Tested and documented

---

## 📚 Documentation

### New Users
1. [README.md](README.md) - Full documentation
2. [QUICKSTART.md](docs/QUICKSTART.md) - 5-minute guide
3. [DEMO.md](DEMO.md) - See it in action

### Contributors
1. [GETTING_STARTED.md](GETTING_STARTED.md) - Dev setup
2. [CONTRIBUTING.md](CONTRIBUTING.md) - How to contribute
3. [PROJECT_OVERVIEW.md](PROJECT_OVERVIEW.md) - Architecture

### Reference
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md) - Command cheat sheet
- [LANGUAGES.md](docs/LANGUAGES.md) - Supported languages
- [EXAMPLES.md](docs/EXAMPLES.md) - Usage examples
- [DOCUMENTATION_INDEX.md](DOCUMENTATION_INDEX.md) - All docs

---

## 🎯 Common Tasks

### Run a program
```bash
polyglot run hello.py
```

### Pass arguments
```bash
polyglot run script.py --args arg1,arg2,arg3
```

### Check toolchains
```bash
polyglot check
```

### List languages
```bash
polyglot list
```

### Get help
```bash
polyglot --help
```

---

## 🛠️ Development

### Build
```bash
make build
```

### Test
```bash
make test
```

### Add a language
See [GETTING_STARTED.md](GETTING_STARTED.md) - takes ~15 minutes!

---

## 📊 Project Status

**Version**: 0.1.0  
**Milestone**: 1 (MVP) - ✅ COMPLETE  
**Languages**: 5 supported  
**Documentation**: 17 files  
**Status**: Production Ready

See [STATUS.md](STATUS.md) for details.

---

## 🗺️ Roadmap

- ✅ **Milestone 1**: MVP (5 languages) - COMPLETE
- 📋 **Milestone 2**: Compiled languages (Java, C, C++, Rust)
- 📋 **Milestone 3**: 20+ languages
- 📋 **Milestone 4**: Esoteric languages
- 📋 **Milestone 5**: Configuration system

See [ROADMAP.md](ROADMAP.md) for full plan.

---

## 🤝 Contributing

We welcome contributions!

1. Read [GETTING_STARTED.md](GETTING_STARTED.md)
2. Check [CONTRIBUTING.md](CONTRIBUTING.md)
3. Pick an issue or add a language
4. Submit a PR

Adding a new language takes ~15 minutes!

---

## 📄 License

MIT License - See [LICENSE](LICENSE)

---

## 🔗 Links

- **Repository**: https://github.com/yourusername/polyglot
- **Issues**: https://github.com/yourusername/polyglot/issues
- **Documentation**: [DOCUMENTATION_INDEX.md](DOCUMENTATION_INDEX.md)

---

## 💡 Quick Tips

- Use `-q` for quiet mode (great for piping)
- Use `-v` for verbose mode (debugging)
- Run `polyglot check` to see what's installed
- Create an alias: `alias pg='polyglot run'`

---

## 🎉 What's New

### v0.1.0 (Milestone 1)
- ✅ 5 languages supported
- ✅ Full CLI with 5 commands
- ✅ Comprehensive documentation
- ✅ Production ready

See [CHANGELOG.md](CHANGELOG.md) for details.

---

## 📞 Need Help?

- 📖 Read the [Quick Start](docs/QUICKSTART.md)
- 💬 Open an [Issue](https://github.com/yourusername/polyglot/issues)
- 📚 Check [Examples](docs/EXAMPLES.md)
- 🎯 See the [Demo](DEMO.md)

---

## ⭐ Star Us!

If you find Polyglot useful, give us a star on GitHub!

---

**One command to run them all!** 🚀

**Start here**: [README.md](README.md) or [QUICKSTART.md](docs/QUICKSTART.md)
