CREATE TABLE blocks (
    id UUID PRIMARY KEY,

    user_id UUID NOT NULL,

    blocked_user_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_block_self
        CHECK (user_id <> blocked_user_id),

    CONSTRAINT uq_block
        UNIQUE(user_id, blocked_user_id)
);

CREATE INDEX idx_blocks_user
ON blocks(user_id);

CREATE INDEX idx_blocks_blocked_user
ON blocks(blocked_user_id);