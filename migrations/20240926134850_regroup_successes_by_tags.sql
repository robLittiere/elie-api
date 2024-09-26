-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
UPDATE tags SET name = 'PlayGames' WHERE name = 'PlayQuizTag';
UPDATE tags SET name = 'WinGames' WHERE name = 'WonQuizTag';
UPDATE tags SET name = 'Connection' WHERE name = 'ConnectionTag';
UPDATE tags SET name = 'Level' WHERE name = 'LevelTag';
UPDATE tags SET name = 'Avatar' WHERE name = 'AvatarTag';

UPDATE successes SET tag_id = 1 WHERE tag_id = 2;
UPDATE successes SET tag_id = 3 WHERE tag_id = 4;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
UPDATE tags SET name = 'PlayQuizTag' WHERE name = 'PlayGames';
UPDATE tags SET name = 'WonQuizTag' WHERE name = 'WinGames';
UPDATE tags SET name = 'ConnectionTag' WHERE name = 'Connection';
UPDATE tags SET name = 'LevelTag' WHERE name = 'Level';
UPDATE tags SET name = 'AvatarTag' WHERE name = 'Avatar';

UPDATE successes SET tag_id = 2 WHERE tag_id = 1;
UPDATE successes SET tag_id = 4 WHERE tag_id = 3;

-- +goose StatementEnd
