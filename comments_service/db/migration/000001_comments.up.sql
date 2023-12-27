CREATE TABLE "comments" (
    "id" SERIAL PRIMARY KEY,
    "post_id" INT NOT NULL,
    "user_unique_id" UUID NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX comments_post_id_idx ON comments (post_id);