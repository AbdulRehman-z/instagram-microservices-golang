CREATE TABLE "followers" (
    "id" SERIAL PRIMARY KEY,
    "leader_unique_id" uuid NOT NULL,  
    "follower_unique_id" uuid NOT NULL,  
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),

    CHECK (leader_unique_id <> follower_unique_id),
    UNIQUE (leader_unique_id, follower_unique_id)
);

CREATE INDEX followers_leader_unique_id_idx ON followers(leader_unique_id);
CREATE INDEX followers_follower_unique_id_idx ON followers(follower_unique_id);
