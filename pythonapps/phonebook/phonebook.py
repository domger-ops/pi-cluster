import sqlite3
import json

# Connecting to DB
conn = sqlite3.connect("contacts.db")
cur = con.cursor()
