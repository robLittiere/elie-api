-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DELETE FROM tags WHERE name = 'PlayGameTag';
DELETE FROM tags WHERE name = 'WonGameTag';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
INSERT INTO tags (name) VALUES ('PlayGameTag');
INSERT INTO tags (name) VALUES ('WonGameTag');
-- +goose StatementEnd
