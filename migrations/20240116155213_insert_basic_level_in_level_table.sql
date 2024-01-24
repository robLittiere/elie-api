-- +goose Up
-- +goose StatementBegin
    INSERT INTO levels (name, next_level_xp_requirement, currency_won) values ('Basic', 100, 100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM users WHERE level_id = (SELECT id FROM levels WHERE name = 'Basic');
    DELETE FROM levels WHERE name = 'Basic';
-- +goose StatementEnd
