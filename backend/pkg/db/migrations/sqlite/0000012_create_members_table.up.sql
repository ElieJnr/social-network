CREATE TABLE IF NOT EXISTS Membership (
    userId TEXT NOT NULL,
    groupId TEXT NOT NULL,
    role VARCHAR(50),
    joinDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userId, groupId),
    FOREIGN KEY (userId) REFERENCES Users(id),
    FOREIGN KEY (groupId) REFERENCES Groups(id)
);
