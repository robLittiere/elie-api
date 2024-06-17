-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'ConnectionTag'), 'Se connecter', 10, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'ConnectionTag'), 'Se connecter', 30, 'intermediate', 0, 7, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'ConnectionTag'), 'Se connecter', 100, 'advanced', 0, 30, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'PlayQuizTag'), 'Faire un quiz', 10, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'PlayQuizTag'), 'Faire des quizs', 30, 'intermediate', 0, 10, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'PlayQuizTag'), 'Faire des quizs', 100, 'advanced', 0, 20, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'WonQuizTag'), 'Réussir un quiz', 10, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'WonQuizTag'), 'Réussir des quizs', 30, 'intermediate', 0, 10, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'WonQuizTag'), 'Réussir des quizs', 100, 'advanced', 0, 20, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'PlayGameTag'), 'Jouer à des jeux', 30, 'intermediate', 0, 5, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'WonGameTag'), 'Réussir des jeux', 30, 'intermediate', 0, 5, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'LevelTag'), 'Débloquer le niveau suivant', 30, 'intermediate', 0, 1, NOW(), NOW());
INSERT INTO quests (tag_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM tag WHERE name = 'AvatarTag'), 'Débloquer un nouvel avatar', 100, 'advanced', 0, 1, NOW(), NOW());
-- +goose Down
-- +goose StatementBegin
DELETE
FROM quests
WHERE name = 'Se connecter';
DELETE
FROM quests
WHERE name = 'Faire un quiz';
DELETE
FROM quests
WHERE name = 'Faire des quizs';
DELETE
FROM quests
WHERE name = 'Réussir un quiz';
DELETE
FROM quests
WHERE name = 'Réussir des quizs';
DELETE
FROM quests
WHERE name = 'Jouer à des jeux';
DELETE
FROM quests
WHERE name = 'Réussir des jeux';
DELETE
FROM quests
WHERE name = 'Débloquer le niveau suivant';
DELETE
FROM quests
WHERE name = 'Débloquer un nouvel avatar';
-- +goose StatementEnd
