CREATE TABLE author (
    id VARCHAR(200) PRIMARY KEY NOT NULL,
    firstname VARCHAR(200) NOT NULL,
    secondname VARCHAR(200) NOT NULL,
    email VARCHAR(200) UNIQUE,
    age INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE book_category (
    id VARCHAR(200) PRIMARY KEY NOT NULL,
    category_name VARCHAR(200) UNIQUE NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE book (
    id VARCHAR(200) PRIMARY KEY NOT NULL,
    category_id VARCHAR(200) NOT NULL,
    author_id VARCHAR(200) NOT NULL,
    name VARCHAR(200) NOT NULL,
    price REAL NOT NULL,
    definition VARCHAR(200),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
