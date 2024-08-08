CREATE TABLE IF NOT EXISTS Chat_groups (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    members TEXT -- JSON string or comma-separated user IDs
);
