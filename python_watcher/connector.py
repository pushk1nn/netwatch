import sqlite3
con = sqlite3.connect("tutorial.db")

cur = con.cursor()

def query_from_timestamp(timestamp: int) -> int:
    res = cur.execute("SELECT name FROM sqlite_master")
    return res

res.fetchone()
