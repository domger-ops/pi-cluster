import sqlite3
import json

# Connecting to DB
conn = sqlite3.connect("contacts.db")
cur = con.cursor()

# Data example
example = {
    "contacts": [
        {
            "name": "myName",
            "phone": "555-5555"
            "email": "thisPerson@email.com"
        }
    ]
}

# Check if table exists
cur.execute("SELECT name FROM sqlite_master WHERE type='table' AND name='contacts'")
table_exists = cur.fetchone()

if not table_exists:
    # create table for db with columns name, phone and email
    cur.execute('''CREATE TABLE contacts (
                        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
                        "name" TEXT NOT NULL,
                        "phone" TEXT,
                        "email" TEXT
                    )''')

    # Insert example data
    for contact in example["contacts"]:
        cur.execute('''INSERT INTO contacts (name, phone, email) VALUES (?, ?, ?)''',
                    (contact["name"], contact["phone"], contact["email"]))

    conn.commit()