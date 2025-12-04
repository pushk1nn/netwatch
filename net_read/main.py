#!/usr/bin/env python3
import argparse
import os
from connector import check_session
from pathlib import Path

#Grabs Mactimes of file passed to the program
def get_timestamp() -> list:
    path = argparse.ArgumentParser(description="Get a file's atime Unix timestamp.")
    path.add_argument("file", type=Path, help="Path to the file")
    args = path.parse_args()

    file: Path = args.file
    if not file.exists():
        raise SystemExit(f"Not found: {file}")
    if not file.is_file():
        raise SystemExit(f"Not a regular file: {file}")

    stat = os.stat(file, follow_symlinks=False)

    mac_times: list = [int(stat.st_atime), int(stat.st_mtime), int(stat.st_ctime)]

    return mac_times

#Driver code to pass timestamps and receive associated IPs
def main() -> int:
    timestamps: list = get_timestamp()
    for stamp in timestamps:
        ret: str = check_session('../net_watcher/data.sqlite', stamp)
        print(ret)
    return 0

if __name__ == "__main__":
    raise SystemExit(main())
