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
        ('Khvicha', 'Kvaratskhelia', 25, 1000000, 'Forward', 1);  -- Player1 for Napoli

    -- Liverpool #2
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Alisson', 'Becker', 29, 7000000, 'Goalkeeper', 2),
    ('Trent', 'Alexander-Arnold', 23, 6000000, 'Defender', 2),
    ('Virgil', 'van Dijk', 30, 8000000, 'Defender', 2),
    ('Andrew', 'Robertson', 27, 7000000, 'Defender', 2),
    ('Joe', 'Gomez', 24, 5000000, 'Defender', 2),
    ('Fabinho', '#', 28, 6000000, 'Midfielder', 2),
    ('Jordan', 'Henderson', 31, 6000000, 'Midfielder', 2),
    ('Thiago', 'Alcântara', 30, 7000000, 'Midfielder', 2),
    ('Mohamed', 'Salah', 29, 10000000, 'Forward', 2),
    ('Sadio', 'Mané', 29, 9000000, 'Forward', 2),
    ('Roberto', 'Firmino', 30, 8000000, 'Forward', 2);

    -- Manchester City #3
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Ederson', 'Moraes', 28, 8000000, 'Goalkeeper', 3),
    ('João', 'Cancelo', 27, 5000000, 'Defender', 3),
    ('Rúben', 'Dias', 24, 10000000, 'Defender', 3),
    ('Aymeric', 'Laporte', 27, 7000000, 'Defender', 3),
    ('Kyle', 'Walker', 31, 6000000, 'Defender', 3),
    ('Rodri', '', 25, 9000000, 'Midfielder', 3),
    ('Kevin', 'De Bruyne', 30, 11000000, 'Midfielder', 3),
    ('Ilkay', 'Gündogan', 31, 8000000, 'Midfielder', 3),
    ('Raheem', 'Sterling', 27, 10000000, 'Forward', 3),
    ('Gabriel', 'Jesus', 24, 9000000, 'Forward', 3),
    ('Ferran', 'Torres', 21, 7000000, 'Forward', 3);

    -- Real Madrid #4
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Thibaut', 'Courtois', 29, 8000000, 'Goalkeeper', 4),
    ('Eder', 'Militao', 24, 25000000, 'Defender', 4),
    ('David', 'Alaba', 29, 40000000, 'Defender', 4),
    ('Raphaël', 'Varane', 28, 60000000, 'Defender', 4),
    ('Ferland', 'Mendy', 26, 35000000, 'Defender', 4),
    ('Casemiro', '', 29, 70000000, 'Midfielder', 4),
    ('Luka', 'Modric', 36, 15000000, 'Midfielder', 4),
    ('Toni', 'Kroos', 32, 25000000, 'Midfielder', 4),
    ('Vinícius', 'Júnior', 21, 60000000, 'Forward', 4),
    ('Karim', 'Benzema', 34, 75000000, 'Forward', 4),
    ('Marco', 'Asensio', 26, 40000000, 'Forward', 4);

    --Atletico Madrid #5
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Jan', 'Oblak', 29, 80000000, 'Goalkeeper', 5),  -- Jan Oblak for Atletico Madrid
    ('Stefan', 'Savić', 31, 35000000, 'Defender', 5),  -- Stefan Savić for Atletico Madrid
    ('José', 'Giménez', 27, 60000000, 'Defender', 5),  -- José Giménez for Atletico Madrid
    ('Renan', 'Lodi', 23, 45000000, 'Defender', 5),  -- Renan Lodi for Atletico Madrid
    ('Kieran', 'Trippier', 31, 30000000, 'Defender', 5),  -- Kieran Trippier for Atletico Madrid
    ('Koke', '', 29, 60000000, 'Midfielder', 5),  -- Koke for Atletico Madrid
    ('Saúl', 'Ñíguez', 27, 70000000, 'Midfielder', 5),  -- Saúl Ñíguez for Atletico Madrid
    ('Thomas', 'Partey', 28, 50000000, 'Midfielder', 5),  -- Thomas Partey for Atletico Madrid
    ('João', 'Félix', 22, 120000000, 'Forward', 5),  -- João Félix for Atletico Madrid
    ('Luis', 'Suárez', 35, 50000000, 'Forward', 5),  -- Luis Suárez for Atletico Madrid
    ('Ángel', 'Correa', 27, 60000000, 'Forward', 5);  -- Ángel Correa for Atletico Madrid

    -- FC Bayern Munich #6
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Manuel', 'Neuer', 35, 40000000, 'Goalkeeper', 6),  -- Manuel Neuer for FC Bayern Munich
    ('Benjamin', 'Pavard', 26, 25000000, 'Defender', 6),  -- Benjamin Pavard for FC Bayern Munich
    ('Niklas', 'Süle', 26, 40000000, 'Defender', 6),  -- Niklas Süle for FC Bayern Munich
    ('Lucas', 'Hernández', 25, 60000000, 'Defender', 6),  -- Lucas Hernández for FC Bayern Munich
    ('David', 'Alaba', 29, 50000000, 'Defender', 6),  -- David Alaba for FC Bayern Munich
    ('Leon', 'Goretzka', 27, 70000000, 'Midfielder', 6),  -- Leon Goretzka for FC Bayern Munich
    ('Joshua', 'Kimmich', 27, 80000000, 'Midfielder', 6),  -- Joshua Kimmich for FC Bayern Munich
    ('Marc', 'Roca', 25, 20000000, 'Midfielder', 6),  -- Marc Roca for FC Bayern Munich
    ('Thomas', 'Müller', 32, 30000000, 'Forward', 6),  -- Thomas Müller for FC Bayern Munich
    ('Leroy', 'Sané', 25, 90000000, 'Forward', 6),  -- Leroy Sané for FC Bayern Munich
    ('Robert', 'Lewandowski', 33, 100000000, 'Forward', 6);  -- Robert Lewandowski for FC Bayern Munich

    -- Inter Milan #7
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Samir', 'Handanović', 37, 15000000, 'Goalkeeper', 7),  -- Samir Handanović for Inter Milan
    ('Stefan', 'De Vrij', 29, 35000000, 'Defender', 7),  -- Stefan De Vrij for Inter Milan
    ('Alessandro', 'Bastoni', 22, 45000000, 'Defender', 7),  -- Alessandro Bastoni for Inter Milan
    ('Milan', 'Škriniar', 27, 50000000, 'Defender', 7),  -- Milan Škriniar for Inter Milan
    ('Marcelo', 'Brozović', 29, 60000000, 'Midfielder', 7),  -- Marcelo Brozović for Inter Milan
    ('Nicolò', 'Barella', 24, 70000000, 'Midfielder', 7),  -- Nicolò Barella for Inter Milan
    ('Achraf', 'Hakimi', 23, 80000000, 'Midfielder', 7),  -- Achraf Hakimi for Inter Milan
    ('Christian', 'Eriksen', 29, 45000000, 'Midfielder', 7),  -- Christian Eriksen for Inter Milan
    ('Lautaro', 'Martínez', 24, 90000000, 'Forward', 7),  -- Lautaro Martínez for Inter Milan
    ('Edin', 'Džeko', 35, 20000000, 'Forward', 7),  -- Edin Džeko for Inter Milan
    ('Ivan', 'Perišić', 32, 25000000, 'Forward', 7);  -- Ivan Perišić for Inter Milan

    -- FC Barcelona #8
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Marc-André', 'ter Stegen', 29, 60000000, 'Goalkeeper', 8),  -- Marc-André ter Stegen for FC Barcelona
    ('Gerard', 'Piqué', 35, 20000000, 'Defender', 8),  -- Gerard Piqué for FC Barcelona
    ('Ronald', 'Araújo', 22, 30000000, 'Defender', 8),  -- Ronald Araújo for FC Barcelona
    ('Jordi', 'Alba', 33, 35000000, 'Defender', 8),  -- Jordi Alba for FC Barcelona
    ('Sergiño', 'Dest', 21, 40000000, 'Defender', 8),  -- Sergiño Dest for FC Barcelona
    ('Sergio', 'Busquets', 33, 25000000, 'Midfielder', 8),  -- Sergio Busquets for FC Barcelona
    ('Frenkie', 'de Jong', 24, 70000000, 'Midfielder', 8),  -- Frenkie de Jong for FC Barcelona
    ('Pedri', 'González', 19, 80000000, 'Midfielder', 8),  -- Pedri González for FC Barcelona
    ('Lionel', 'Messi', 34, 150000000, 'Forward', 8),  -- Lionel Messi for FC Barcelona
    ('Memphis', 'Depay', 28, 60000000, 'Forward', 8),  -- Memphis Depay for FC Barcelona
    ('Antoine', 'Griezmann', 30, 80000000, 'Forward', 8);  -- Antoine Griezmann for FC Barcelona

    --  Tottenham #9
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Hugo', 'Lloris', 35, 20000000, 'Goalkeeper', 9),
    ('Sergio', 'Reguilón', 25, 35000000, 'Defender', 9),
    ('Eric', 'Dier', 28, 25000000, 'Defender', 9),
    ('Toby', 'Alderweireld', 32, 30000000, 'Defender', 9),
    ('Matt', 'Doherty', 30, 20000000, 'Defender', 9),
    ('Pierre-Emile', 'Højbjerg', 26, 45000000, 'Midfielder', 9),
    ('Giovani', 'Lo Celso', 26, 40000000, 'Midfielder', 9),
    ('Tanguy', 'Ndombele', 25, 50000000, 'Midfielder', 9),
    ('Steven', 'Bergwijn', 24, 35000000, 'Forward', 9),
    ('Harry', 'Kane', 28, 100000000, 'Forward', 9),
    ('Heung-Min', 'Son', 29, 80000000, 'Forward', 9);

--  Chelsea #10
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Edouard', 'Mendy', 29, 25000000, 'Goalkeeper', 10),
    ('Reece', 'James', 22, 50000000, 'Defender', 10),
    ('Thiago', 'Silva', 37, 10000000, 'Defender', 10),
    ('Antonio', 'Rüdiger', 28, 40000000, 'Defender', 10),
    ('Ben', 'Chilwell', 24, 60000000, 'Defender', 10),
    ('NGolo', 'Kanté', 30, 80000000, 'Midfielder', 10),
    ('Jorginho', NULL, 30, 50000000, 'Midfielder', 10),
    ('Mason', 'Mount', 23, 70000000, 'Midfielder', 10),
    ('Christian', 'Pulisic', 23, 60000000, 'Forward', 10),
    ('Timo', 'Werner', 25, 80000000, 'Forward', 10),
    ('Kai', 'Havertz', 22, 80000000, 'Forward', 10);

-- Borussia Dortmund #11
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Roman', 'Bürki', 31, 15000000, 'Goalkeeper', 11),
    ('Mats', 'Hummels', 33, 20000000, 'Defender', 11),
    ('Manuel', 'Akanji', 26, 30000000, 'Defender', 11),
    ('Raphael', 'Guerreiro', 28, 35000000, 'Defender', 11),
    ('Thomas', 'Meunier', 30, 25000000, 'Defender', 11),
    ('Jude', 'Bellingham', 18, 60000000, 'Midfielder', 11),
    ('Axel', 'Witsel', 33, 15000000, 'Midfielder', 11),
    ('Giovanni', 'Reyna', 19, 40000000, 'Midfielder', 11),
    ('Marco', 'Reus', 32, 30000000, 'Forward', 11),
    ('Erling', 'Haaland', 21, 100000000, 'Forward', 11),
    ('Jadon', 'Sancho', 22, 80000000, 'Forward', 11);

    -- PSG #12
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Keylor', 'Navas', 35, 15000000, 'Goalkeeper', 12),
    ('Marquinhos', NULL, 27, 60000000, 'Defender', 12),
    ('Presnel', 'Kimpembe', 26, 40000000, 'Defender', 12),
    ('Achraf', 'Hakimi', 23, 60000000, 'Defender', 12),
    ('Layvin', 'Kurzawa', 29, 20000000, 'Defender', 12),
    ('Georginio', 'Wijnaldum', 31, 60000000, 'Midfielder', 12),
    ('Marco', 'Verratti', 29, 80000000, 'Midfielder', 12),
    ('Idrissa', 'Gueye', 32, 30000000, 'Midfielder', 12),
    ('Neymar', NULL, 30, 120000000, 'Forward', 12),
    ('Kylian', 'Mbappé', 23, 150000000, 'Forward', 12),
    ('Lionel', 'Messi', 34, 150000000, 'Forward', 12);

    -- Juventus #13
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Wojciech', 'Szczęsny', 31, 20000000, 'Goalkeeper', 13),
    ('Giorgio', 'Chiellini', 37, 10000000, 'Defender', 13),
    ('Leonardo', 'Bonucci', 34, 15000000, 'Defender', 13),
    ('Matthijs', 'de Ligt', 22, 75000000, 'Defender', 13),
    ('Danilo', NULL, 30, 25000000, 'Defender', 13),
    ('Weston', 'McKennie', 23, 40000000, 'Midfielder', 13),
    ('Rodrigo', 'Bentancur', 24, 50000000, 'Midfielder', 13),
    ('Arthur', NULL, 25, 60000000, 'Midfielder', 13),
    ('Federico', 'Chiesa', 24, 80000000, 'Forward', 13),
    ('Paulo', 'Dybala', 28, 90000000, 'Forward', 13),
    ('Alvaro', 'Morata', 29, 60000000, 'Forward', 13);

    -- AC Milan #14
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Gianluigi', 'Donnarumma', 23, 60000000, 'Goalkeeper', 14),
    ('Theo', 'Hernández', 24, 70000000, 'Defender', 14),
    ('Fikayo', 'Tomori', 24, 35000000, 'Defender', 14),
    ('Simon', 'Kjær', 32, 10000000, 'Defender', 14),
    ('Davide', 'Calabria', 25, 25000000, 'Defender', 14),
    ('Franck', 'Kessié', 25, 50000000, 'Midfielder', 14),
    ('Sandro', 'Tonali', 22, 40000000, 'Midfielder', 14),
    ('Ismaël', 'Bennacer', 24, 50000000, 'Midfielder', 14),
    ('Ante', 'Rebić', 28, 60000000, 'Forward', 14),
    ('Zlatan', 'Ibrahimović', 40, 5000000, 'Forward', 14),
    ('Rafael', 'Leão', 23, 70000000, 'Forward', 14);

    -- Manchester United #15
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('David', 'De Gea', 31, 35000000, 'Goalkeeper', 15),
    ('Harry', 'Maguire', 29, 80000000, 'Defender', 15),
    ('Luke', 'Shaw', 26, 50000000, 'Defender', 15),
    ('Aaron', 'Wan-Bissaka', 24, 60000000, 'Defender', 15),
    ('Raphaël', 'Varane', 28, 70000000, 'Defender', 15),
    ('Paul', 'Pogba', 29, 90000000, 'Midfielder', 15),
    ('Bruno', 'Fernandes', 27, 80000000, 'Midfielder', 15),
    ('Scott', 'McTominay', 25, 40000000, 'Midfielder', 15),
    ('Marcus', 'Rashford', 24, 100000000, 'Forward', 15),
    ('Jadon', 'Sancho', 22, 90000000, 'Forward', 15),
    ('Cristiano', 'Ronaldo', 37, 25000000, 'Forward', 15);

    -- Arsenal #16
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Bernd', 'Leno', 30, 30000000, 'Goalkeeper', 16),
    ('Kieran', 'Tierney', 24, 50000000, 'Defender', 16),
    ('Gabriel', 'Magalhães', 24, 45000000, 'Defender', 16),
    ('Ben', 'White', 24, 50000000, 'Defender', 16),
    ('Héctor', 'Bellerín', 26, 35000000, 'Defender', 16),
    ('Thomas', 'Partey', 28, 60000000, 'Midfielder', 16),
    ('Granit', 'Xhaka', 29, 40000000, 'Midfielder', 16),
    ('Emile', 'Smith Rowe', 21, 60000000, 'Midfielder', 16),
    ('Bukayo', 'Saka', 20, 80000000, 'Forward', 16),
    ('Pierre-Emerick', 'Aubameyang', 33, 50000000, 'Forward', 16),
    ('Alexandre', 'Lacazette', 30, 40000000, 'Forward', 16);







