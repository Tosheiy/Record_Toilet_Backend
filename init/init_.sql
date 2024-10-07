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

/*
-- MySQL用
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=1049to -d -v 
/Users/toshi/Document/dev_app/Record_Toilet/Record_Toilet_Backend/model/init.sql:
/docker-entrypoint-initdb.d/init.sql mysql:latest

-- deploy用
mysql -h mytoiletrecord.cpwiqm8ec2pr.ap-northeast-1.rds.amazonaws.com -u admin -p
*/

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