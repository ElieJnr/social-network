CREATE TABLE IF NOT EXISTS sessions(
    sessionId TEXT PRIMARY KEY,
    userId TEXT,
    username TEXT NOT NULL,
    expired_at TEXT NOT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);