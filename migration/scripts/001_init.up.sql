-- 001_init.up.sql
-- users
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) NOT NULL UNIQUE,
    name VARCHAR(255),
    username VARCHAR(100) NOT NULL UNIQUE,
    age INT,
    pass VARCHAR(255),
    salt VARCHAR(255),
    permission VARCHAR(50) DEFAULT 'user',
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    INDEX idx_username (username)
    );

-- electricity_used
CREATE TABLE IF NOT EXISTS electricity_useds (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) NOT NULL UNIQUE,
    date_used DATETIME(3),
    username VARCHAR(100),
    elec_used DOUBLE,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    INDEX idx_username (username),
    CONSTRAINT fk_electricity_user FOREIGN KEY (username)
    REFERENCES users(username)
    ON UPDATE CASCADE
    ON DELETE CASCADE
    );
