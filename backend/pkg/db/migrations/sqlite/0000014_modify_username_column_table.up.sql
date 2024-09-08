-- Migration pour modifier la colonne 'username' dans la table 'Users'

-- Crée une nouvelle table avec la structure modifiée
CREATE TABLE IF NOT EXISTS Users_new (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    username TEXT,  -- Modifié, suppression de UNIQUE
    dateOfBirth TEXT NOT NULL,
    bio TEXT,
    avatar TEXT,
    isPrivate BOOLEAN NOT NULL
);

-- Copie les données de l'ancienne table vers la nouvelle
INSERT INTO Users_new (id, email, password, firstname, lastname, username, dateOfBirth, bio, avatar, isPrivate)
SELECT id, email, password, firstname, lastname, username, dateOfBirth, bio, avatar, isPrivate
FROM Users;

-- Supprime l'ancienne table
DROP TABLE Users;

-- Renomme la nouvelle table pour utiliser l'ancien nom
ALTER TABLE Users_new RENAME TO Users;
