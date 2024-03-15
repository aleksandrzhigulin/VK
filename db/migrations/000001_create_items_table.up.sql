CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance INT
);

CREATE TABLE IF NOT EXISTS quests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    cost INT
);

CREATE TABLE IF NOT EXISTS quests_users (
    user_id INT,
    quest_id INT,
    PRIMARY KEY (user_id, quest_id),
    CONSTRAINT vk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT vk_quest FOREIGN KEY (quest_id) REFERENCES quests(id)
);