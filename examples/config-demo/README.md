# Configuration Demo

This example demonstrates Polyglot's configuration system.

## Files

- `.polyglot.yaml` - Project configuration
- `hello.py` - Python script
- `hello.js` - JavaScript script
- `custom.pyx` - Custom extension (Cython-like)
- `app.jsx` - React JSX file

## Configuration Features Demonstrated

### 1. Custom Extensions

The `.polyglot.yaml` maps custom extensions:
- `.pyx` → Python (for Cython files)
- `.jsx` → JavaScript (for React)

### 2. Language-Specific Settings

Python is configured with:
- Optimization flag (`-O`)
- Custom environment variable

JavaScript is configured with:
- Development environment

### 3. Environment Variables

Global environment variables are set for all executions.

## Usage

```bash
# Run Python file
polyglot run hello.py

# Run JavaScript file
polyglot run hello.js

# Run custom extension (treated as Python)
polyglot run custom.pyx

# Run JSX file (treated as JavaScript)
polyglot run app.jsx

# View configuration
polyglot config show

# Get specific values
polyglot config get languages.python.version
polyglot config get custom_extensions..pyx
```

## Expected Behavior

All files should run successfully, with custom extensions automatically mapped to the correct language.
