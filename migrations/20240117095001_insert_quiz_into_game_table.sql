-- +goose Up
-- +goose StatementBegin
    INSERT INTO games (name, description, catch_phrase, can_be_multiplayer, created_at, updated_at) VALUES ('Quiz', 'A simple quiz game', 'Répondez à des quizzs surprenants pour en apprendre plus sur l''écologie', true, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM games WHERE name = 'Quiz';
-- +goose StatementEnd
