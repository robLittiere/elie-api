-- +goose Up
-- +goose StatementBegin
    INSERT INTO quest_types (type) VALUES ('daily');
    INSERT INTO quest_types (type) VALUES ('weekly');
    INSERT INTO quest_types (type) VALUES ('monthly');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM quest_types WHERE type = 'daily';
    DELETE FROM quest_types WHERE type = 'weekly';
    DELETE FROM quest_types WHERE type = 'monthly';
-- +goose StatementEnd
