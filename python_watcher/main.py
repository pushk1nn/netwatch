#!/usr/bin/env python3
import argparse
import os
from pathlib import Path

def get_timestamp() -> int:
    path = argparse.ArgumentParser(description="Get a file's atime Unix timestamp.")
    path.add_argument("file", type=Path, help="Path to the file")
    args = path.parse_args()

    file: Path = args.file
    if not file.exists():
        raise SystemExit(f"Not found: {file}")
    if not file.is_file():
        raise SystemExit(f"Not a regular file: {file}")

    st = os.stat(file, follow_symlinks=False)

    atime_timestamp = int(st.st_atime)
    return atime_timestamp


def main() -> int:
    timestamp: int = get_timestamp()
    print(timestamp)
    return 0

if __name__ == "__main__":
    raise SystemExit(main())
