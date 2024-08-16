CREATE TABLE IF NOT EXISTS UserPost (
    id TEXT PRIMARY KEY,
    postId TEXT NOT NULL,
    userId TEXT NOT NULL,
    statut TEXT NOT NULL,
    FOREIGN KEY (postid) REFERENCES Posts(id),
    FOREIGN KEY (userid) REFERENCES Users(id)
);
