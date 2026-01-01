#!/usr/bin/env python3
"""Check for redundant comments in Go code."""

import re
import sys
from pathlib import Path

REDUNDANT_PATTERNS = [
    (r'// (\w+) (\w+)s? (the )?\1', 'Comment repeats function name'),
    (r'// Get\w+ gets? the \w+', 'Obvious getter comment'),
    (r'// Set\w+ sets? the \w+', 'Obvious setter comment'),
    (r'// New\w+ creates? a new \w+', 'Obvious constructor comment'),
    (r'// \w+ takes .* and returns', 'Comment just describes signature'),
    (r'^// \w+$', 'Single-word comment (not useful)'),
]

def check_file(filepath):
    """Check a single Go file for redundant comments."""
    errors = []

    with open(filepath, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    for i, line in enumerate(lines, start=1):
        stripped = line.strip()

        if not stripped.startswith('//'):
            continue

        for pattern, message in REDUNDANT_PATTERNS:
            if re.search(pattern, stripped, re.IGNORECASE):
                errors.append(f"{filepath}:{i}: {message}: {stripped}")

    return errors

def main(files):
    """Check all provided files."""
    all_errors = []

    for filepath in files:
        if Path(filepath).suffix == '.go':
            errors = check_file(filepath)
            all_errors.extend(errors)

    if all_errors:
        print("ERROR: Redundant comments detected:")
        for error in all_errors:
            print(f"  {error}")
        print("\nRemove redundant comments. Document only what's not obvious from code.")
        sys.exit(1)

    sys.exit(0)

if __name__ == '__main__':
    main(sys.argv[1:])
