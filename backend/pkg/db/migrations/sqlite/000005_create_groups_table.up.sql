CREATE TABLE IF NOT EXISTS Chat_groups (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    userid INTEGER NOT NULL,
    description TEXT,
    FOREIGN KEY (userid) REFERENCES Users(id)
);
