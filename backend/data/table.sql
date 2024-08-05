CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    bio TEXT,
    avatar TEXT,
    isPrivate BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS Posts (
    id INTEGER PRIMARY KEY,
    userId INTEGER NOT NULL,
    content TEXT NOT NULL,
    category TEXT,
    creatdate DATETIME DEFAULT CURRENT_TIMESTAMP,
    isPrivate BOOLEAN NOT NULL,
    FOREIGN KEY (userId) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Comments (
    id INTEGER PRIMARY KEY,
    postid INTEGER NOT NULL,
    userid INTEGER NOT NULL,
    content TEXT NOT NULL,
    creatdate DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (postid) REFERENCES Posts(id),
    FOREIGN KEY (userid) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Chats (
    id INTEGER PRIMARY KEY,
    senderId INTEGER NOT NULL,
    receverId INTEGER NOT NULL,
    content TEXT NOT NULL,
    sendAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (senderId) REFERENCES Users(id),
    FOREIGN KEY (receverId) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Groups (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    members TEXT -- JSON string or comma-separated user IDs
);

CREATE TABLE IF NOT EXISTS Events (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    dateTime DATETIME NOT NULL,
    response TEXT
);

CREATE TABLE IF NOT EXISTS Followers (
    id INTEGER PRIMARY KEY,
    followerid INTEGER NOT NULL,
    followingid INTEGER NOT NULL,
    status TEXT NOT NULL,
    requestTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (followerid) REFERENCES Users(id),
    FOREIGN KEY (followingid) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Profil (
    id INTEGER PRIMARY KEY,
    userActivity TEXT,
    followingUser TEXT, -- JSON string or comma-separated user IDs
    followerUser TEXT, -- JSON string or comma-separated user IDs
    isPrivate BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS Notifications (
    id INTEGER PRIMARY KEY,
    senderId INTEGER NOT NULL,
    receverId INTEGER NOT NULL,
    content TEXT NOT NULL,
    sendAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (senderId) REFERENCES Users(id),
    FOREIGN KEY (receverId) REFERENCES Users(id)
);
