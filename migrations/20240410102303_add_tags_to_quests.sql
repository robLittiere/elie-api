-- +goose Up
-- +goose StatementBegin
ALTER TABLE quests ADD COLUMN IF NOT EXISTS tags varchar(255) DEFAULT NULL;
UPDATE quests SET tags = 'quiz,game' WHERE name LIKE 'Faire un quiz%';
UPDATE quests SET tags = 'quiz_complete,game' WHERE name LIKE '%Réussir des quizs%';
UPDATE quests SET tags = 'quiz_complete,game' WHERE name LIKE '%Réussir un quiz%';
UPDATE quests SET tags = 'game' WHERE name LIKE '%Jouer à des jeux%';
UPDATE quests SET tags = 'login' WHERE name LIKE '%Se connecter%';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quests DROP COLUMN IF EXISTS tags;
-- +goose StatementEnd
