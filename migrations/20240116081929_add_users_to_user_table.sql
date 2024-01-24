-- +goose Up
-- +goose StatementBegin
INSERT INTO users (uuid, level_id, email, password, username, xp, currency_amount)
VALUES ('c0bf4924-e82b-4bcb-8a06-31443d3a7f4e', (SELECT id FROM levels WHERE level_number = 1), 'rob@mail.com', 'rob', 'robinou', 0, 0);
INSERT INTO users (uuid, level_id, email, password, username, xp, currency_amount)
VALUES ('c0bf4924-e82b-4bcb-8a06-31443d3a7f4f', (SELECT id FROM levels WHERE level_number = 2), 'test2@mail.com', 'rob', 'robinouxxx', 1000, 50);
INSERT INTO users (uuid, level_id, email, password, username, xp, currency_amount)
VALUES ('c0bf4924-e82b-4bcb-8a06-31443d3a7f4g', (SELECT id FROM levels WHERE level_number = 3), 'barbara@mail.com', 'barbara', 'barbara', 100, 200);
INSERT INTO users (uuid, level_id, email, password, username, xp, currency_amount)
VALUES ('c0bf4924-e82b-4bcb-8a06-31443d3a7f4h', (SELECT id FROM levels WHERE level_number = 2), 'lena@mail.com', 'lena', 'lena', 200, 650);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE username = 'robinou';
DELETE FROM users WHERE username = 'robinouxxx';
DELETE FROM users WHERE username = 'barbara';
DELETE FROM users WHERE username = 'lena';
-- +goose StatementEnd
