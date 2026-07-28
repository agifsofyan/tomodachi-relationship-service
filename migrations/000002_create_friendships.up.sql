CREATE TABLE friendships (
    id UUID PRIMARY KEY,

    user_id UUID NOT NULL,

    friend_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_friendship_self
        CHECK (user_id <> friend_id),

    CONSTRAINT uq_friendship
        UNIQUE (user_id, friend_id)
);

CREATE INDEX idx_friendships_user
ON friendships(user_id);

CREATE INDEX idx_friendships_friend
ON friendships(friend_id);