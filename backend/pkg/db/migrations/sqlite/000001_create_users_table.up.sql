CREATE TABLE IF NOT EXISTS Users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    username TEXT,
    dateOfBirth TEXT NOT NULL,
    bio TEXT,
    avatar TEXT,
    isPrivate BOOLEAN NOT NULL
);
