CREATE TABLE toilet_records (
    id INTEGER PRIMARY KEY,
    description TEXT,
    created_at DATETIME,
    length INTEGER,
    location TEXT,
    feeling INTEGER,
    uid TEXT
);

CREATE TABLE user_table (
    id INTEGER PRIMARY KEY,
    utid VARCHAR(32) UNIQUE NOT NULL,
    uid TEXT UNIQUE NOT NULL,
    apikey VARCHAR(50) UNIQUE NOT NULL
);