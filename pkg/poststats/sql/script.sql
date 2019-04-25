CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE post_views (
    post_uid UUID PRIMARY KEY,
    num_views INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE post_votes (
    post_uid UUID NOT NULL,
    user_uid UUID NOT NULL,
    vote INTEGER NOT NULL DEFAULT 0,
    CONSTRAINT post_votes_pk PRIMARY KEY (post_uid, user_uid)
);

