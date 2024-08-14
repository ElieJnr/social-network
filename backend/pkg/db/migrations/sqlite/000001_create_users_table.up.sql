CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    username TEXT UNIQUE,
    date_of_birth TEXT NOT NULL,
    bio TEXT,
    avatar TEXT,
    session TEXT,
    isPrivate BOOLEAN NOT NULL
);
