#!/usr/bin/env python3
import sys

print(f"Arguments received: {len(sys.argv) - 1}")
for i, arg in enumerate(sys.argv[1:], 1):
    print(f"  arg{i}: {arg}")
