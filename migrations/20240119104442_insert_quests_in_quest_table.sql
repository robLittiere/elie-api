-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'daily'), 'Se connecter', 0, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'daily'), 'Faire un quiz', 5, 'easy', 0, 1, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'daily'), 'Réussir un quiz', 10, 'intermediate', 0, 1, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'daily'), 'Réussir des quizs', 20, 'advanced', 0, 3, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'weekly'), 'Se connecter', 10, 'easy', 0, 3, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'weekly'), 'Jouer à des jeux', 20, 'intermediate', 0, 5, NOW(), NOW());
INSERT INTO quests (quest_type_id, name, xp, difficulty, currency_reward, done_condition, created_at, updated_at)
VALUES ((SELECT id FROM quest_types WHERE type = 'weekly'), 'Réussir des quizs', 50, 'advanced', 0, 5, NOW(), NOW());
-- +goose Down
-- +goose StatementBegin
DELETE
FROM user_quests
WHERE quest_id IN (SELECT id FROM quests WHERE name = 'Se connecter');
DELETE
FROM quests
WHERE name = 'Se connecter';
DELETE
FROM user_quests
WHERE quest_id IN (SELECT id FROM quests WHERE name = 'Faire un quiz');
DELETE
FROM quests
WHERE name = 'Faire un quiz';
DELETE
FROM user_quests
WHERE quest_id IN (SELECT id FROM quests WHERE name = 'Réussir un quiz');
DELETE
FROM quests
WHERE name = 'Réussir un quiz';
DELETE
FROM user_quests
WHERE quest_id IN (SELECT id FROM quests WHERE name = 'Réussir des quizs');
DELETE
FROM quests
WHERE name = 'Réussir des quizs';
DELETE
FROM user_quests
WHERE quest_id IN (SELECT id FROM quests WHERE name = 'Jouer à des jeux');
DELETE
FROM quests
WHERE name = 'Jouer à des jeux';
-- +goose StatementEnd
