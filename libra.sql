CREATE TABLE IF NOT EXISTS users (
    share_key uuid UNIQUE PRIMARY KEY,
    username TEXT UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    fullname TEXT,
    bio TEXT,
    whoareyou TEXT,
    avatar BLOB,
    account_rating REAL,
	post_rating REAL,
	follower_count INTEGER,
	following_count INTEGER,
	post_count INTEGER,
	location text,
    session_id uuid UNIQUE,
    birthday TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    created_at TEXT NOT NULL,
    active boolean DEFAULT true
);

CREATE TABLE IF NOT EXISTS followers (
    follower TEXT,
    followed TEXT,
    created_at TEXT,
    FOREIGN KEY (follower) REFERENCES users (share_key),
    FOREIGN KEY (followed) REFERENCES users (share_key)
);

CREATE TABLE post (
    id INT,
    user_id INT,
    title VARCHAR,
    content VARCHAR,
     updated_at TEXT,
    created_at TEXT
);

CREATE TABLE comment (
    id INT,
    user_id INT,
    post_id INT,
    content VARCHAR,
    updated_at TEXT,
    created_at TEXT
);

CREATE TABLE like (
    id INT,
    user_id INT,
    post_id INT,
    created_at TEXT
);

CREATE TABLE message (
    id INT,
    sender_id INT,
    receiver_id INT,
    content VARCHAR,
    created_at TEXT
);