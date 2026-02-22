#!/usr/bin/env python3
"""Demo Python script for configuration testing."""

import os

print("Hello from Python!")
print(f"Environment: {os.getenv('NODE_ENV', 'not set')}")
print(f"Demo var: {os.getenv('DEMO', 'not set')}")
print("✓ Python configuration working!")
