CREATE TABLE "accounts" (
    id serial PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    unique_id VARCHAR(255) UNIQUE NOT NULL,
    age INT NOT NULL,
    bio TEXT,
    avatar VARCHAR(255),
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX account_id_idx ON "accounts"("id");