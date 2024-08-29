CREATE TABLE IF NOT EXISTS Followers (
    id TEXT PRIMARY KEY,
    userId TEXT NOT NULL,
    followedId TEXT NOT NULL,
    statut BOOLEAN,
    FOREIGN KEY (userId) REFERENCES Users(id),
    FOREIGN KEY (followedId) REFERENCES Users(id)
);
