-- +goose Up
-- +goose StatementBegin
    INSERT INTO tags (name) VALUES ('PlayQuizTag');
    INSERT INTO tags (name) VALUES ('PlayGameTag');
    INSERT INTO tags (name) VALUES ('WonQuizTag');
    INSERT INTO tags (name) VALUES ('WonGameTag');
    INSERT INTO tags (name) VALUES ('ConnectionTag');
    INSERT INTO tags (name) VALUES ('LevelTag');
    INSERT INTO tags (name) VALUES ('AvatarTag');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM tags WHERE name = 'PlayQuizTag';
    DELETE FROM tags WHERE name = 'PlayGameTag';
    DELETE FROM tags WHERE name = 'WonQuizTag';
    DELETE FROM tags WHERE name = 'WonGameTag';
    DELETE FROM tags WHERE name = 'ConnectionTag';
    DELETE FROM tags WHERE name = 'LevelTag';
    DELETE FROM tags WHERE name = 'AvatarTag';
-- +goose StatementEnd
