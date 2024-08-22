CREATE TABLE IF NOT EXISTS Followers (
    id TEXT PRIMARY KEY,
    userId TEXT NOT NULL,
    follwedId TEXT NOT NULL,
    FOREIGN KEY (userid) REFERENCES Users(id)
);
