-- +goose Up
-- +goose StatementBegin
INSERT INTO successes (tag_id, name, xp, progression_rank, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tags WHERE name = 'ConnectionTag' LIMIT 1), 'Se connecter', 1000, 1, 100, 100, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'ConnectionTag' LIMIT 1), 'Se connecter', 2000, 2, 100, 200, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'PlayQuizTag' LIMIT 1), 'Faire des quizs', 1000, 1, 100, 100, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'PlayQuizTag' LIMIT 1), 'Faire des quizs', 2000, 2, 100, 200, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'WonQuizTag' LIMIT 1), 'Réussir des quizs', 2000, 1, 200, 100, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'WonQuizTag' LIMIT 1), 'Réussir des quizs', 4000, 2, 200, 200, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'PlayGameTag' LIMIT 1), 'Jouer à des jeux', 1000, 1, 100, 100, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'PlayGameTag' LIMIT 1), 'Jouer à des jeux', 2000, 2, 100, 200, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'WonGameTag' LIMIT 1), 'Gagner des jeux', 2000, 1, 200, 100, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'WonGameTag' LIMIT 1), 'Gagner des jeux', 4000, 2, 200, 200, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'LevelTag' LIMIT 1), 'Monter de niveau', 1000, 1, 100, 1, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'LevelTag' LIMIT 1), 'Monter de niveau', 2000, 2, 100, 2, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'AvatarTag' LIMIT 1), 'Débloquer un nouvel avatar', 1000, 1, 0, 1, NOW(), NOW()),
       ((SELECT id FROM tags WHERE name = 'AvatarTag' LIMIT 1), 'Débloquer un nouvel avatar', 2000, 2, 0, 2, NOW(), NOW());
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
