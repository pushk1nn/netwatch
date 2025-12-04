
import sqlite3
from datetime import datetime

def check_session(database: str, timestamp: int) -> str:
    # Convert the input timestamp to a string format for SQL query
    #timestamp_str = timestamp.timestamp()

    conn = sqlite3.connect(database)
    cursor = conn.cursor()

    query = """
    SELECT c1.ip
    FROM connections c1
    JOIN connections c2 ON c1.ip = c2.ip
        WHERE c1.type = 'CONNECT'
            AND c2.type = 'DISCONNECT'
            AND c1.unix_time <= ?
            AND c2.unix_time >= ?;
    """

    cursor.execute(query, (timestamp, timestamp))
    ret = cursor.fetchone()
    conn.close()

    return ret[0]

# Example Usage
db_path = '../net_watcher/data.sqlite'
timestamp_to_check = datetime.fromisoformat('2025-12-02 21:06:45.129551-05:00').timestamp()

print(check_session(db_path, timestamp_to_check))
