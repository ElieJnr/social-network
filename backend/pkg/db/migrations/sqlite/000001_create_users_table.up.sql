CREATE TABLE IF NOT EXISTS Users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    passwords TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    username TEXT UNIQUE,
    dateOfBirth TEXT NOT NULL,
    bio TEXT,
    avatar TEXT,
    isPrivate BOOLEAN NOT NULL
);
