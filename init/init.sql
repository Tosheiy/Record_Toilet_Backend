CREATE DATABASE IF NOT EXISTS RecordToilet;
USE RecordToilet;
CREATE TABLE IF NOT EXISTS toilet_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT,
    created_at DATETIME,
    length INT,
    location TEXT,
    feeling INT,
    uid VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS user_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    utid VARCHAR(32) UNIQUE NOT NULL,
    uid VARCHAR(255) UNIQUE NOT NULL,
    apikey VARCHAR(50) UNIQUE NOT NULL
);