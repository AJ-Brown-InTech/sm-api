CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    user_password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    bio TEXT,
    avatar BYTEA,
    account_rating REAL,
    real_rating_number REAL,
    rating_average REAL,
    follower_count INTEGER DEFAULT 0,
    following_count INTEGER DEFAULT 0,
    location POINT,
    birthday TEXT NOT NULL,
    city TEXT,
    country TEXT,
    state_province TEXT,
    updated_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    active BOOLEAN DEFAULT true, 
    verified BOOLEAN DEFAULT false
);

CREATE TABLE IF NOT EXISTS followers (
    follower UUID,
    followed UUID,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (follower) REFERENCES users (user_id),
    FOREIGN KEY (followed) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS posts (
    post_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (user_id),
    caption VARCHAR,
    photo BYTEA,
    photo_post BOOLEAN,
    text_post BOOLEAN,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (user_id),
    post_id SERIAL REFERENCES posts (post_id),
    caption VARCHAR,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS postrating ( 
    post_rating_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (user_id),
    post_id SERIAL REFERENCES posts (post_id),
    post_rating REAL,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS messages (
    message_id SERIAL PRIMARY KEY,
    sender_id UUID REFERENCES users (user_id),
    receiver_id UUID REFERENCES users (user_id),
    content VARCHAR,
    photo BYTEA,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ratings (
    rating_id SERIAL PRIMARY KEY,
    rater_id UUID REFERENCES users (user_id),
    rated_id UUID REFERENCES users (user_id),
    rating_value REAL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tags (
    tag_id SERIAL PRIMARY KEY,
    tag_name TEXT
);

CREATE TABLE IF NOT EXISTS post_tags (
    post_id SERIAL REFERENCES posts (post_id),
    tag_id SERIAL REFERENCES tags (tag_id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (post_id, tag_id)
);

CREATE TABLE IF NOT EXISTS notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (user_id),
    notification_type TEXT,
    notification_content TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
