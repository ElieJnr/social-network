CREATE TABLE IF NOT EXISTS Members (
    id INTEGER PRIMARY KEY,
    groupid INTEGER NOT NULL,
    userid INTEGER NOT NULL,
    FOREIGN KEY (groupid) REFERENCES Chat_groups(id),
    FOREIGN KEY (userid) REFERENCES Users(id)
);
