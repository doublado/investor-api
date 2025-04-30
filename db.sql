CREATE DATABASE IF NOT EXISTS investor CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE investor;

CREATE TABLE IF NOT EXISTS leaderboard (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL UNIQUE,
    score INT NOT NULL
);

INSERT INTO leaderboard (name, score)
VALUES ('Nikolaj', 100);