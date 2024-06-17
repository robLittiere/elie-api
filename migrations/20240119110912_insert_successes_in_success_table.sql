-- +goose Up
-- +goose StatementBegin
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'ConnectionTag'), 'Se connecter', 0, 'easy',1, 100, 10, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'ConnectionTag'), 'Se connecter', 0, 'intermediate', 2, 200, 20, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'ConnectionTag'), 'Se connecter', 0, 'advanced', 3, 400, 30, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'WonQuizTag'), 'Réussir des quizs', 30, 'easy', 1, 100, 10, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'WonQuizTag'), 'Réussir des quizs', 60, 'intermediate', 2, 200, 20, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'WonQuizTag'), 'Réussir des quizs', 100, 'advanced', 3, 500, 50, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'LevelTag'), 'Atteindre le niveau 2', 0, 'easy', 1, 0, 1, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'LevelTag'), 'Atteindre le niveau 3', 0, 'easy', 2, 0, 1, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'LevelTag'), 'Atteindre le niveau 5', 0, 'intermediate', 3, 100, 1, NOW(), NOW());
    INSERT INTO successes (tag_id, name, xp, difficulty, progression_rank, currency_reward, done_condition, created_at, updated_at)
    VALUES ((SELECT id FROM tag WHERE name = 'AvatarTag'), 'Débloquer un nouvel avatar', 100, 'intermediate', 1, 0, 1, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM successes WHERE name = 'Se connecter';
    DELETE FROM successes WHERE name = 'Réussir des quizs';
    DELETE FROM successes WHERE name = 'Atteindre le niveau 2';
    DELETE FROM successes WHERE name = 'Atteindre le niveau 3';
    DELETE FROM successes WHERE name = 'Atteindre le niveau 5';
    DELETE FROM successes WHERE name = 'Débloquer un nouvel avatar';
-- +goose StatementEnd
