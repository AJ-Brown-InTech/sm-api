CREATE TABLE IF NOT EXISTS users (
    user_id uuid PRIMARY KEY,
    username TEXT UNIQUE not null unique,
    email TEXT NOT NULL UNIQUE,
    user_password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    bio TEXT,
    avatar BYTEA,
    account_rating REAL,
    real_rating_number REAL,
    rating_average REAL,
	follower_count INTEGER default 0,
	following_count INTEGER default 0,
	location point,
    birthday TEXT NOT NULL,
    city TEXT,
    country TEXT,
    state_province TEXT,
    updated_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    active boolean DEFAULT true, 
    verified boolean DEFAULT false
);

CREATE TABLE IF NOT EXISTS followers (
    follower TEXT,
    followed TEXT,
    created_at TEXT,
    FOREIGN KEY (follower) REFERENCES users (user_id) ON DELETE CASC,
    FOREIGN KEY (followed) REFERENCES users (user_id) ON DELETE CASC
);

CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (user_id),
    caption VARCHAR,
    photo BYTEA,
    photo_post BOOLEAN,
    text_post BOOLEAN,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
     user_id UUID REFERENCES users (user_id),
    post_id SERIAL REFERENCES posts (post_id),
    caption VARCHAR,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW()
    )


CREATE TABLE post_rating ( -- equivalent to a likes tb ininstagram
    post_rating_id serial primary key,
    user_id UUID REFERENCES users (user_id),
    post_id SERIAL REFERENCES posts (post_id),
    post_rating REAL,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE messages (
    message_id SERIAL PRIMARY KEY,
    sender_id UUID REFERENCES users (user_id),
    receiver_id UUID REFERENCES users (user_id),
    content VARCHAR,
    photo BYTEA,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE ratings (
    rating_id SERIAL PRIMARY KEY,
    rater_id UUID REFERENCES users (user_id),
    rated_id UUID REFERENCES users (user_id),
    rating_value REAL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE tags (
    tag_id SERIAL PRIMARY KEY,
    tag_name TEXT
);

CREATE TABLE post_tags (
    post_id SERIAL REFERENCES posts (post_id),
    tag_id SERIAL REFERENCES tags (tag_id),
    PRIMARY KEY (post_id, tag_id)
);

CREATE TABLE notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (user_id),
    notification_type TEXT,
    notification_content TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
