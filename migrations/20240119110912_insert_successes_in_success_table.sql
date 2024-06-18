-- +goose Up
-- +goose StatementBegin
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'ConnectionTag'), 'Se connecter', 1000, 1, 100, 100, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'PlayQuizTag'), 'Faire des quizs', 1000, 1, 100, 100, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'WonQuizTag'), 'Réussir des quizs', 2000, 1, 200, 100, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'PlayGameTag'), 'Jouer à des jeux', 1000, 1, 100, 100, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'WonGameTag'), 'Gagner des jeux', 2000, 1, 200, 100, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'LevelTag'), 'Monter de niveau', 1000, 1, 100, 1, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tags WHERE name = 'AvatarTag'), 'Débloquer un nouvel avatar', 1000, 1, 0, 1, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM successes WHERE name = 'Se connecter';
    DELETE FROM successes WHERE name = 'Faire des quizs';
    DELETE FROM successes WHERE name = 'Réussir des quizs';
    DELETE FROM successes WHERE name = 'Jouer à des jeux';
    DELETE FROM successes WHERE name = 'Gagner des jeux';
    DELETE FROM successes WHERE name = 'Monter de niveau';
    DELETE FROM successes WHERE name = 'Débloquer un nouvel avatar';
-- +goose StatementEnd
