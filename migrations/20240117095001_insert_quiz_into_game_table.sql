-- +goose Up
-- +goose StatementBegin
    INSERT INTO games (name, description, catch_phrase, can_be_multiplayer, game_version, created_at, updated_at) VALUES ('Quiz', 'A simple quiz game', 'Répondez à des quizzs surprenants pour en apprendre plus sur l''écologie', true, '1.0', NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM quiz_games where gid = (SELECT id FROM games WHERE name = 'Quiz');
    DELETE FROM games WHERE name = 'Quiz';
-- +goose StatementEnd
