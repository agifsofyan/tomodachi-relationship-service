CREATE TABLE friend_requests (
    id UUID PRIMARY KEY,

    requester_id UUID NOT NULL,
    receiver_id UUID NOT NULL,

    status VARCHAR(20) NOT NULL,

    responded_at TIMESTAMPTZ NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_friend_request_self
        CHECK (requester_id <> receiver_id),

    CONSTRAINT uq_friend_request
        UNIQUE (requester_id, receiver_id)
);

CREATE INDEX idx_friend_requests_requester
ON friend_requests(requester_id);

CREATE INDEX idx_friend_requests_receiver
ON friend_requests(receiver_id);

CREATE INDEX idx_friend_requests_status
ON friend_requests(status);