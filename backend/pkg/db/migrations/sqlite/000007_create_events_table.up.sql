CREATE TABLE IF NOT EXISTS Events (
    id TEXT PRIMARY KEY,                 
    memberId INTEGER NOT NULL,           
    groupId INTEGER NOT NULL,           
    title TEXT NOT NULL,                 
    description TEXT NOT NULL,           
    eventDate TEXT NOT NULL,         
    FOREIGN KEY (groupId) REFERENCES Groups(id)  
);

CREATE TABLE IF NOT EXISTS EventResponses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique identifier for the response
    eventId TEXT NOT NULL,                 -- ID of the event
    memberId INTEGER NOT NULL,             -- ID of the member responding
    response TEXT NOT NULL,                -- User's response (e.g., 'Going', 'Not Going')
    FOREIGN KEY (eventId) REFERENCES Events(id),   -- Foreign key reference to the Events table
    FOREIGN KEY (memberId) REFERENCES Members(id)  -- Foreign key reference to the Members table
);
