-- +goose Up
-- +goose StatementBegin
INSERT INTO levels (name, next_level_xp_requirement, currency_won) VALUES ('beginner', 200, 100);
INSERT INTO levels (name, next_level_xp_requirement, currency_won) VALUES ('intermediate', 400, 200);
INSERT INTO levels (name, next_level_xp_requirement, currency_won) VALUES ('advanced', 800, 400);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM levels WHERE name = 'beginner';
DELETE FROM levels WHERE name = 'intermediate';
DELETE FROM levels WHERE name = 'advanced';
-- +goose StatementEnd
