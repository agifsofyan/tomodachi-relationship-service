CREATE UNIQUE INDEX idx_friendships_user_friend
ON friendships(user_id, friend_id);