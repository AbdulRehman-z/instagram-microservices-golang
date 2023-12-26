CREATE TABLE post_likes (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    user_unique_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(post_id, user_unique_id)
); 

CREATE TABLE comment_likes (
    id SERIAL PRIMARY KEY,
    comment_id INT NOT NULL,
    user_unique_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(comment_id, user_unique_id)
);
