CREATE TABLE Coach (
    coach_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    exp_year INTEGER
);

-- Add a team_id column to the Coach table
ALTER TABLE Coach
ADD COLUMN team_id INTEGER;

-- Add a foreign key constraint to connect team_id in Coach to the Teams table
ALTER TABLE Coach
ADD CONSTRAINT fk_coach_team
FOREIGN KEY (team_id) REFERENCES Teams(team_id);


CREATE TABLE Teams (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(100),
    coach_id INTEGER,
    FOREIGN KEY (coach_id) REFERENCES Coach(coach_id)
);

CREATE TABLE Player (
    player_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    player_age INTEGER,
    player_cost DECIMAL,
    player_pos VARCHAR(50),
    team_id INTEGER,
    FOREIGN KEY (team_id) REFERENCES Teams(team_id)
);



-- Coach insertion
INSERT INTO Coach (first_name, last_name, exp_year, team_id)
VALUES
    ('Carlo', 'Ancelotti', 15, 1),          -- Napoli
    ('Jürgen', 'Klopp', 20, 2),             -- Liverpool
    ('Pep', 'Guardiola', 10, 3),             -- Manchester City
    ('Zinedine', 'Zidane', 10, 4),          -- Real Madrid
    ('Diego', 'Simeone', 12, 5),            -- Atletico Madrid
    ('Julian', 'Nagelsmann', 7, 6),         -- FC Bayern Munich
    ('Simone', 'Inzaghi', 6, 7),            -- Inter Milan
    ('Xavi', 'Hernandez', 4, 8),            -- FC Barcelona
    ('Antonio', 'Conte', 18, 9),            -- Tottenham Hotspur
    ('Thomas', 'Tuchel', 12, 10),           -- Chelsea
    ('Marco', 'Rose', 5, 11),               -- Borussia Dortmund
    ('Mauricio', 'Pochettino', 10, 12),     -- PSG
    ('Massimiliano', 'Allegri', 10, 13),    -- Juventus
    ('Stefano', 'Pioli', 5, 14),            -- AC Milan
    ('Ralf', 'Rangnick', 5, 15),            -- Manchester United
    ('Mikel', 'Arteta', 6, 16);             -- Arsenal


-- Teams insertion
-- Teams insertion
INSERT INTO Teams (team_name, coach_id)
VALUES
    ('Napoli', 1),
    ('Liverpool', 2),
    ('Manchester City', 3),
    ('Real Madrid', 4),
    ('Atlectico Madrid', 5),
    ('FC Bayern Munich', 6),
    ('Inter Milan', 7),
    ('FC Barcelona', 8),
    ('Tottenham Hotspur', 9),
    ('Chelsea', 10),
    ('Borussia Dortmund', 11),
    ('PSG', 12),
    ('Juventus', 13),
    ('AC Milan', 14),
    ('Manchester United', 15),
    ('Arsenal', 16);

-- PLayer insertion
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    -- Napoli #1
        ('Alex', 'Meret', 23, 800000, 'Goalkeeper', 1),  -- Player3 for Napoli
        ('Juan', 'Jesus', 26, 1200000, 'Defender', 1),  -- Player4 for Napoli
        ('Giovanni', 'Di Lorenzo', 26, 1200000, 'Defender', 1),  -- Player4 for Napoli
        ('Mathias', 'Olivera', 26, 1200000, 'Defender', 1),  -- Player4 for Napoli
        ('Amir', 'Rrahmani', 26, 1200000, 'Defender', 1),  -- Player4 for Napoli
        ('Jesper', 'Lindstorm', 27, 1500000, 'Midfielder', 1),  -- Player2 for Napoli
        ('Diego', 'Demme', 27, 1500000, 'Midfielder', 1),  -- Player2 for Napoli
        ('Stanislav', 'Lobotka', 27, 1500000, 'Midfielder', 1),  -- Player2 for Napoli
        ('Giovanni', 'Simeone', 25, 1000000, 'Forward', 1),   -- Player1 for Napoli
        ('Giacomo', 'Raspadori', 25, 1000000, 'Forward', 1),   -- Player1 for Napoli
        ('Khvicha', 'Kvaratskhelia', 25, 1000000, 'Forward', 1)  -- Player1 for Napoli



