CREATE TABLE
  IF NOT EXISTS users (
    id integer primary key autoincrement not null,
    username TEXT unique,
    email TEXT not null unique,
    password TEXT not null,
    fullname TEXT,
    bio TEXT,
    occupation TEXT,
    avatar BLOB,
    rating TEXT,
    city TEXT,
    state TEXT,
    session_id TEXT unique,
    birthday TIMESTAMP not null,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
  );
