CREATE TABLE IF NOT EXISTS Notifications (
    id TEXT PRIMARY KEY,
    user_receiver_id TEXT NOT NULL,  -- Référence à l'utilisateur qui reçoit la notification
    user_id TEXT NOT NULL       -- Référence à l'utilisateur qui envoie la notification
    type TEXT NOT NULL,        -- Type de notification (par exemple, 'follow_request', 'group_invitation', 'event_created', etc.)
    message TEXT,              -- Message de notification ou détails
    is_read BOOLEAN DEFAULT FALSE, -- Indique si la notification a été lue
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Date et heure de la création
    FOREIGN KEY (user_receiver_id) REFERENCES users(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
