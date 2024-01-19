-- +goose Up
-- +goose StatementBegin
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'rob@mail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND quest_type_id = (SELECT id FROM quest_types WHERE type = 'daily')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'rob@mail.com'), (SELECT id FROM quests WHERE name = 'Jouer à des jeux' AND quest_type_id = (SELECT id FROM quest_types WHERE type = 'weekly')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'lena@mail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND quest_type_id = (SELECT id FROM quest_types WHERE type = 'daily')), 0, false, NOW(), NOW());
    INSERT INTO user_quests (user_id, quest_id, progression, is_completed, created_at, updated_at) VALUES ((SELECT id  from users WHERE email = 'lena@mail.com'), (SELECT id FROM quests WHERE name = 'Réussir des quizs' AND quest_type_id = (SELECT id FROM quest_types WHERE type = 'weekly')), 0, false, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DELETE FROM user_quests WHERE user_id = (SELECT id  from users WHERE email = 'rob@mail.com');
    DELETE FROM user_quests WHERE user_id = (SELECT id  from users WHERE email = 'lena@mail.com');
-- +goose StatementEnd
