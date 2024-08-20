CREATE TABLE IF NOT EXISTS Members (
    id TEXT PRIMARY KEY,
    groupid INTEGER NOT NULL,
    userId TEXT NOT NULL,
    FOREIGN KEY (groupid) REFERENCES Chat_groups(id),
    FOREIGN KEY (userid) REFERENCES Users(id)
);
