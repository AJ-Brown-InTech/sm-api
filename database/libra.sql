CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    fullname TEXT,
    bio TEXT,
    occupation TEXT,
    avatar BLOB,
    rating REAL,
    city TEXT,
    state TEXT,
    session_id TEXT UNIQUE,
    birthday TEXT NOT NULL,
    updated_at TEXT,
    created_at TEXT
);
