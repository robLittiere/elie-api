-- +goose Up
-- +goose StatementBegin
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'rob@mail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND difficulty = 'easy' AND tag_id = (SELECT id FROM tag WHERE name = 'WonQuizTag')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'rob@mail.com'), (SELECT id FROM quests WHERE name = 'Jouer à des jeux' AND difficulty = 'easy' AND tag_id = (SELECT id FROM tag WHERE name = 'PlayGameTag')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'lena@mail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND difficulty = 'easy' AND tag_id = (SELECT id FROM tag WHERE name = 'WonQuizTag')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'lena@mail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND difficulty = 'intermediate' AND tag_id = (SELECT id FROM tag WHERE name = 'WonQuizTag')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'bbr@gmail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND difficulty = 'advanced' AND tag_id = (SELECT id FROM tag WHERE name = 'WonQuizTag')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'bbr@gmail.com'), (SELECT id FROM quests WHERE name = 'Faire des quizs' AND difficulty = 'intermediate' AND tag_id = (SELECT id FROM tag WHERE name = 'PlayQuizTag')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'bbr@gmail.com'), (SELECT id FROM quests WHERE name = 'Se connecter' AND difficulty = 'easy' AND tag_id = (SELECT id FROM tag WHERE name = 'ConnectionTag')), 0, false, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM user_quests WHERE user_id = (SELECT id  from users WHERE email = 'rob@mail.com');
    DELETE FROM user_quests WHERE user_id = (SELECT id  from users WHERE email = 'lena@mail.com');
    DELETE FROM user_quests WHERE user_id = (SELECT id  from users WHERE email = 'bbr@gmail.com');
-- +goose StatementEnd
