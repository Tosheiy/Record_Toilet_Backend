-- CREATE TABLE toilet_records (
--     id INTEGER PRIMARY KEY,
--     description TEXT,
--     created_at DATETIME,
--     length INTEGER,
--     location TEXT,
--     feeling INTEGER,
--     uid TEXT
-- );

-- CREATE TABLE user_table (
--     id INTEGER PRIMARY KEY,
--     utid VARCHAR(32) UNIQUE NOT NULL,
--     uid TEXT UNIQUE NOT NULL,
--     apikey VARCHAR(50) UNIQUE NOT NULL
-- );

-- MySQL用
CREATE TABLE toilet_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT,
    created_at DATETIME,
    length INT,
    location TEXT,
    feeling INT,
    uid VARCHAR(255)
);

CREATE TABLE user_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    utid VARCHAR(32) UNIQUE NOT NULL,
    uid VARCHAR(255) UNIQUE NOT NULL,
    apikey VARCHAR(50) UNIQUE NOT NULL
);