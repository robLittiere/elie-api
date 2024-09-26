-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE successes ADD COLUMN short_name VARCHAR(255);
UPDATE successes SET short_name = 'Login' WHERE name = 'Se connecter';
UPDATE successes SET short_name = 'LevelUp' WHERE name = 'Monter de niveau';
UPDATE successes SET short_name = 'UnlockAvatar' WHERE name = 'Débloquer un nouvel avatar';
UPDATE successes SET short_name = 'PlayGames' WHERE name = 'Jouer à des jeux';
UPDATE successes SET short_name = 'WinGames' WHERE name = 'Gagner des jeux';
UPDATE successes SET short_name = 'PlayQuizzes' WHERE name = 'Faire des quizs';
UPDATE successes SET short_name = 'WinQuizzes' WHERE name = 'Réussir des quizs';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE successes DROP COLUMN short_name;
-- +goose StatementEnd
