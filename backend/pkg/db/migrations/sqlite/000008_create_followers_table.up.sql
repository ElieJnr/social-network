CREATE TABLE IF NOT EXISTS Followers (
    id TEXT PRIMARY KEY,
    userid TEXT NOT NULL,
    follwerId TEXT NOT NULL,
    FOREIGN KEY (userid) REFERENCES Users(id)
);
