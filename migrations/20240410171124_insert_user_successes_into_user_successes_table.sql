-- +goose Up
-- +goose StatementBegin
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'rob@mail.com'), (SELECT id FROM successes WHERE name = 'Se connecter'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'rob@mail.com'), (SELECT id FROM successes WHERE name = 'Faire des quizs'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'lena@mail.com'), (SELECT id FROM successes WHERE name = 'Réussir des quizs'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'lena@mail.com'), (SELECT id FROM successes WHERE name = 'Jouer à des jeux'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Se connecter'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Faire des quizs'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Réussir des quizs'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Jouer à des jeux'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Gagner des jeux'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Monter de niveau'), 0, false, NOW(), NOW());
INSERT INTO user_successes (user_id, success_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'barbara@mail.com'), (SELECT id FROM successes WHERE name = 'Débloquer un nouvel avatar'), 0, false, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_successes WHERE user_id = (SELECT id  from users WHERE email = 'rob@mail.com');
DELETE FROM user_successes WHERE user_id = (SELECT id  from users WHERE email = 'lena@mail.com');
DELETE FROM user_successes WHERE user_id = (SELECT id  from users WHERE email = 'barbara@mail.com');
-- +goose StatementEnd