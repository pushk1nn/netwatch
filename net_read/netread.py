#!/usr/bin/env python3
import argparse
import os
from connector import check_session
from pathlib import Path
import datetime
import pytz

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

def output_time(res):
    if res:
        for ip in res:
            print(f"Connection from {ip}")

#Driver code to pass timestamps and receive associated IPs
def main() -> int:
    timestamps: list = get_timestamp()
    # for stamp in timestamps:
    #     ret: str = check_session('../net_watcher/data.sqlite', stamp)
    #     print(ret)

    est_timezone = pytz.timezone('America/New_York')

    res = check_session('../net_watcher/data.sqlite', timestamps[0])
    print("ACCESSED:", datetime.datetime.fromtimestamp(timestamps[0], tz=est_timezone))
    output_time(res)

    res = check_session('../net_watcher/data.sqlite', timestamps[1])
    print("\nMODIFIED:", datetime.datetime.fromtimestamp(timestamps[1], tz=est_timezone))
    output_time(res)

    res = check_session('../net_watcher/data.sqlite', timestamps[2])
    print("\nCHANGED:", datetime.datetime.fromtimestamp(timestamps[2], tz=est_timezone))
    output_time(res)


    return 0

if __name__ == "__main__":
    raise SystemExit(main())
