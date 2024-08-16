CREATE TABLE IF NOT EXISTS Followers (
    id TEXT PRIMARY KEY,
    follwerId TEXT NOT NULL,
    FOREIGN KEY (userid) REFERENCES Users(id)
);
