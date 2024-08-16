CREATE TABLE IF NOT EXISTS chatGroups (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    userId TEXT NOT NULL,
    description TEXT,
    FOREIGN KEY (userid) REFERENCES Users(id)
);
