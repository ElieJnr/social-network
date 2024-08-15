CREATE TABLE IF NOT EXISTS Followers (
    user_id INTEGER PRIMARY KEY,
    followerId INTEGER NOT NULL,
    FOREIGN KEY (followerId) REFERENCES Users(id)
);
