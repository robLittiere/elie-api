-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE successes ADD COLUMN parent_id INT NULL;
UPDATE successes SET parent_id = 7 WHERE name = 'Faire des quizs';
UPDATE successes SET parent_id = 9 WHERE name = 'Réussir des quizs';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE successes DROP COLUMN parent_id;
-- +goose StatementEnd
