CREATE TABLE IF NOT EXISTS LikesDislikes (
    like_dislike_id INTEGER PRIMARY KEY,
    post_id INTEGER,
    comment_id INTEGER,
    user_id INTEGER,
    liked BOOLEAN NOT NULL DEFAULT FALSE,
    disliked BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (post_id) REFERENCES Posts(post_id),
    FOREIGN KEY (comment_id) REFERENCES Comments(comment_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);
