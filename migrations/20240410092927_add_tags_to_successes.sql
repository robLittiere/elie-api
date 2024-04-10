-- +goose Up
-- +goose StatementBegin
ALTER TABLE successes ADD COLUMN IF NOT EXISTS tags varchar(255) DEFAULT NULL;
UPDATE successes SET tags = 'login' WHERE name LIKE '%Se connecter%';
UPDATE successes SET tags = 'quiz' WHERE name LIKE '%quiz%';
UPDATE successes SET tags = 'level' WHERE name LIKE '%niveau%';
UPDATE successes SET tags = 'avatar' WHERE name LIKE '%avatar%';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE successes DROP COLUMN IF EXISTS tags;
-- +goose StatementEnd
