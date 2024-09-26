-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
UPDATE quests SET tag_id = 1 WHERE tag_id = 2;
UPDATE quests SET tag_id = 3 WHERE tag_id = 4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
UPDATE quests SET tag_id = 2 WHERE tag_id = 1;
UPDATE quests SET tag_id = 4 WHERE tag_id = 3;
-- +goose StatementEnd
