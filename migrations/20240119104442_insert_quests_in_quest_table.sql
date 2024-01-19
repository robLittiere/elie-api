-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (1, 'Se connecter', 0, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (1, 'Faire un quiz', 5, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (1, 'Réussir un quiz', 10, 'intermediate', 0, 1, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (1, 'Réussir des quizs', 20, 'advanced', 0, 3, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (2, 'Se connecter', 10, 'easy', 0, 3, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (2, 'Jouer à des jeux', 20, 'intermediate', 0, 5, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES (2, 'Réussir des quizs', 50, 'advanced', 0, 5, NOW(), NOW());
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
WHERE name = 'Réussir un quiz';
DELETE
FROM quests
WHERE name = 'Réussir des quizs';
DELETE
FROM quests
WHERE name = 'Jouer à des jeux';
-- +goose StatementEnd
