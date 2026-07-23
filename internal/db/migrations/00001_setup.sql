-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
    `id` text NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `name` text NOT NULL,
    `fingerprint` text NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE INDEX IF NOT EXISTS `idx_users_created_at` ON `users` (`created_at`);

CREATE TABLE IF NOT EXISTS `sessions` (
    `id` text NOT NULL,
    `user_id` text NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `name` text NOT NULL,
    `visibility` text NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `messages` (
    `id` text NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `size` integer NOT NULL,
    `content` blob NOT NULL,
    `user_id` text NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    `session_id` text NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    PRIMARY KEY (`id`)
);

CREATE INDEX IF NOT EXISTS idx_messages_session_created ON messages (session_id, created_at);
CREATE INDEX IF NOT EXISTS idx_messages_user_created ON messages (user_id, created_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `messages`;
DROP TABLE IF EXISTS `sessions`;
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
