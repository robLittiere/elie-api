-- +goose Up
-- +goose StatementBegin
    INSERT INTO tag (name) VALUES ('PlayQuizTag');
    INSERT INTO tag (name) VALUES ('PlayGameTag');
    INSERT INTO tag (name) VALUES ('WonQuizTag');
    INSERT INTO tag (name) VALUES ('WonGameTag');
    INSERT INTO tag (name) VALUES ('ConnectionTag');
    INSERT INTO tag (name) VALUES ('LevelTag');
    INSERT INTO tag (name) VALUES ('AvatarTag');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM tag WHERE name = 'PlayQuizTag';
    DELETE FROM tag WHERE name = 'PlayGameTag';
    DELETE FROM tag WHERE name = 'WonQuizTag';
    DELETE FROM tag WHERE name = 'WonGameTag';
    DELETE FROM tag WHERE name = 'ConnectionTag';
    DELETE FROM tag WHERE name = 'LevelTag';
    DELETE FROM tag WHERE name = 'AvatarTag';
-- +goose StatementEnd
