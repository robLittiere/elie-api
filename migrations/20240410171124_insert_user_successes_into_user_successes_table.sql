-- +goose Up
-- +goose StatementBegin
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at)
VALUES ((SELECT id FROM users WHERE email = 'rob@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Se connecter' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'rob@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Faire des quizs' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'lena@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Réussir des quizs' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'lena@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Jouer à des jeux' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Se connecter' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Faire des quizs' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Réussir des quizs' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Jouer à des jeux' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Gagner des jeux' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Monter de niveau' LIMIT 1), 0, false, NOW(), NOW()),
       ((SELECT id FROM users WHERE email = 'barbara@mail.com' LIMIT 1), (SELECT id FROM successes WHERE name = 'Débloquer un nouvel avatar' LIMIT 1), 0, false, NOW(), NOW());
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE FROM user_successes WHERE user_id = (SELECT id  from users WHERE email = 'rob@mail.com');
DELETE FROM user_successes WHERE user_id = (SELECT id  from users WHERE email = 'lena@mail.com');
DELETE FROM user_successes WHERE user_id = (SELECT id  from users WHERE email = 'barbara@mail.com');
-- +goose StatementEnd