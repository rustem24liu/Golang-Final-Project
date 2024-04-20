CREATE DATABASE football_team;

CREATE TABLE Teams (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(100)
);

CREATE TABLE Coach (
    coach_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    exp_year INTEGER,
    team_id INTEGER,
    FOREIGN KEY (team_id) REFERENCES Teams(team_id)
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
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    activated BOOLEAN NOT NULL,
    permissions TEXT[]
);


INSERT INTO Teams (team_name)
VALUES
    ('Napoli'),
    ('Liverpool'),
    ('Manchester City'),
    ('Real Madrid'),
    ('Atletico Madrid'),
    ('FC Bayern Munich'),
    ('Inter Milan'),
    ('FC Barcelona'),
    ('Tottenham Hotspur'),
    ('Chelsea'),
    ('Borussia Dortmund'),
    ('PSG'),
    ('Juventus'),
    ('AC Milan'),
    ('Manchester United'),
    ('Arsenal');

-- Coach insertion
INSERT INTO Coach (first_name, last_name, exp_year, team_id)
VALUES
    ('Carlo', 'Ancelotti', 15, 1),
    ('Jürgen', 'Klopp', 20, 2),
    ('Pep', 'Guardiola', 10, 3),
    ('Zinedine', 'Zidane', 10, 4),
    ('Diego', 'Simeone', 12, 5),
    ('Julian', 'Nagelsmann', 7, 6),
    ('Simone', 'Inzaghi', 6, 7),
    ('Xavi', 'Hernandez', 4, 8),
    ('Antonio', 'Conte', 18, 9),
    ('Thomas', 'Tuchel', 12, 10),
    ('Marco', 'Rose', 5, 11),
    ('Mauricio', 'Pochettino', 10, 12),
    ('Massimiliano', 'Allegri', 10, 13),
    ('Stefano', 'Pioli', 5, 14),
    ('Ralf', 'Rangnick', 5, 15),
    ('Mikel', 'Arteta', 6, 16);

-- Team: Manchester City #1
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Ederson', 'Moraes', 29, 80000000, 'Goalkeeper', 1),
    ('Kyle', 'Walker', 32, 50000000, 'Defender', 1),
    ('John', 'Stones', 28, 40000000, 'Defender', 1),
    ('Rúben', 'Dias', 25, 80000000, 'Defender', 1),
    ('João', 'Cancelo', 28, 60000000, 'Defender', 1),
    ('Kevin', 'De Bruyne', 30, 120000000, 'Midfielder', 1),
    ('Rodri', 'Hernandes', 25, 70000000, 'Midfielder', 1),
    ('Phil', 'Foden', 21, 90000000, 'Midfielder', 1),
    ('Riyad', 'Mahrez', 31, 60000000, 'Forward', 1),
    ('Gabriel', 'Jesus', 25, 70000000, 'Forward', 1),
    ('Jack', 'Grealish', 26, 100000000, 'Forward', 1);

-- Team: Real Madrid #2
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Thibaut', 'Courtois', 29, 60000000, 'Goalkeeper', 2),
    ('Dani', 'Carvajal', 30, 40000000, 'Defender', 2),
    ('Éder', 'Militão', 24, 35000000, 'Defender', 2),
    ('Sergio', 'Ramos', 36, 10000000, 'Defender', 2),
    ('Ferland', 'Mendy', 27, 50000000, 'Defender', 2),
    ('Carlos', 'Casemiro', 30, 80000000, 'Midfielder', 2),
    ('Luka', 'Modrić', 36, 15000000, 'Midfielder', 2),
    ('Toni', 'Kroos', 32, 50000000, 'Midfielder', 2),
    ('Vinícius', 'Júnior', 22, 80000000, 'Forward', 2),
    ('Karim', 'Benzema', 34, 80000000, 'Forward', 2),
    ('Eden', 'Hazard', 32, 60000000, 'Forward', 2);

-- Team: Liverpool #3
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Alisson', 'Becker', 29, 80000000, 'Goalkeeper', 3),
    ('Trent', 'Alexander-Arnold', 24, 100000000, 'Defender', 3),
    ('Virgil', 'van Dijk', 30, 90000000, 'Defender', 3),
    ('Joe', 'Gomez', 25, 50000000, 'Defender', 3),
    ('Andrew', 'Robertson', 28, 80000000, 'Defender', 3),
    ('Fabio', 'Tovares', 28, 60000000, 'Midfielder', 3),
    ('Jordan', 'Henderson', 32, 20000000, 'Midfielder', 3),
    ('Thiago', 'Alcântara', 31, 40000000, 'Midfielder', 3),
    ('Sadio', 'Mané', 30, 90000000, 'Forward', 3),
    ('Mohamed', 'Salah', 30, 120000000, 'Forward', 3),
    ('Diogo', 'Jota', 25, 60000000, 'Forward', 3);

-- Team: FC Barcelona #4
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Marc-André', 'ter Stegen', 30, 80000000, 'Goalkeeper', 4),
    ('Sergiño', 'Dest', 22, 50000000, 'Defender', 4),
    ('Ronald', 'Araújo', 23, 70000000, 'Defender', 4),
    ('Gerard', 'Piqué', 35, 20000000, 'Defender', 4),
    ('Jordi', 'Alba', 33, 30000000, 'Defender', 4),
    ('Frenkie', 'de Jong', 25, 100000000, 'Midfielder', 4),
    ('Pedri', 'Lopes', 19, 80000000, 'Midfielder', 4),
    ('Sergio', 'Busquets', 33, 20000000, 'Midfielder', 4),
    ('Ansu', 'Fati', 19, 110000000, 'Forward', 4),
    ('Memphis', 'Depay', 29, 50000000, 'Forward', 4),
    ('Ousmane', 'Dembélé', 25, 60000000, 'Forward', 4);

-- Team: Bayern Munich #5
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Manuel', 'Neuer', 36, 10000000, 'Goalkeeper', 5),
    ('Benjamin', 'Pavard', 26, 40000000, 'Defender', 5),
    ('Niklas', 'Süle', 26, 35000000, 'Defender', 5),
    ('Lucas', 'Hernández', 26, 70000000, 'Defender', 5),
    ('Alphonso', 'Davies', 21, 80000000, 'Defender', 5),
    ('Joshua', 'Kimmich', 27, 100000000, 'Midfielder', 5),
    ('Leon', 'Goretzka', 27, 80000000, 'Midfielder', 5),
    ('Kingsley', 'Coman', 25, 60000000, 'Forward', 5),
    ('Thomas', 'Müller', 32, 25000000, 'Forward', 5),
    ('Leroy', 'Sané', 26, 70000000, 'Forward', 5),
    ('Robert', 'Lewandowski', 34, 100000000, 'Forward', 5);

-- Team: Manchester United #6
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('David', 'De Gea', 31, 35000000, 'Goalkeeper', 6),
    ('Harry', 'Maguire', 29, 80000000, 'Defender', 6),
    ('Luke', 'Shaw', 26, 50000000, 'Defender', 6),
    ('Aaron', 'Wan-Bissaka', 24, 60000000, 'Defender', 6),
    ('Raphaël', 'Varane', 28, 70000000, 'Defender', 6),
    ('Paul', 'Pogba', 29, 90000000, 'Midfielder', 6),
    ('Bruno', 'Fernandes', 27, 80000000, 'Midfielder', 6),
    ('Scott', 'McTominay', 25, 40000000, 'Midfielder', 6),
    ('Marcus', 'Rashford', 24, 100000000, 'Forward', 6),
    ('Jadon', 'Sancho', 22, 90000000, 'Forward', 6),
    ('Cristiano', 'Ronaldo', 37, 25000000, 'Forward', 6);

-- Team: Chelsea #7
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Edouard', 'Mendy', 30, 50000000, 'Goalkeeper', 7),
    ('César', 'Azpilicueta', 33, 20000000, 'Defender', 7),
    ('Antonio', 'Rüdiger', 29, 40000000, 'Defender', 7),
    ('Andreas', 'Christensen', 26, 30000000, 'Defender', 7),
    ('Reece', 'James', 22, 60000000, 'Defender', 7),
    ('N''Golo', 'Kanté', 31, 80000000, 'Midfielder', 7),
    ('Mateo', 'Kovačić', 28, 50000000, 'Midfielder', 7),
    ('Mason', 'Mount', 23, 80000000, 'Midfielder', 7),
    ('Kai', 'Havertz', 23, 90000000, 'Forward', 7),
    ('Timo', 'Werner', 26, 70000000, 'Forward', 7),
    ('Christian', 'Pulisic', 25, 60000000, 'Forward', 7);

-- Team: Paris Saint-Germain #8
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Keylor', 'Navas', 35, 15000000, 'Goalkeeper', 8),
    ('Achraf', 'Hakimi', 23, 70000000, 'Defender', 8),
    ('Presnel', 'Kimpembe', 26, 60000000, 'Defender', 8),
    ('Marqos', 'Correa', 28, 80000000, 'Defender', 8),
    ('Nuno', 'Mendes', 20, 50000000, 'Defender', 8),
    ('Georginio', 'Wijnaldum', 31, 40000000, 'Midfielder', 8),
    ('Marco', 'Verratti', 29, 60000000, 'Midfielder', 8),
    ('Idrissa', 'Gueye', 32, 30000000, 'Midfielder', 8),
    ('Neymar', 'Jr', 31, 150000000, 'Forward', 8),
    ('Kylian', 'Mbappé', 24, 150000000, 'Forward', 8),
    ('Lionel', 'Messi', 35, 25000000, 'Forward', 8);

-- Team: Juventus #9
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Wojciech', 'Szczęsny', 32, 30000000, 'Goalkeeper', 9),
    ('Juan', 'Cuadrado', 34, 20000000, 'Defender', 9),
    ('Matthijs', 'de Ligt', 23, 80000000, 'Defender', 9),
    ('Giorgio', 'Chiellini', 38, 5000000, 'Defender', 9),
    ('Alex', 'Sandro', 31, 20000000, 'Defender', 9),
    ('Arthur', 'Melo', 26, 50000000, 'Midfielder', 9),
    ('Rodrigo', 'Bentancur', 25, 40000000, 'Midfielder', 9),
    ('Weston', 'McKennie', 24, 30000000, 'Midfielder', 9),
    ('Federico', 'Chiesa', 24, 80000000, 'Forward', 9),
    ('Paulo', 'Dybala', 29, 70000000, 'Forward', 9),
    ('Álvaro', 'Morata', 30, 40000000, 'Forward', 9);

-- Team: Sevilla #10
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Yassine', 'Bounou', 31, 20000000, 'Goalkeeper', 10),
    ('Jesús', 'Navas', 36, 5000000, 'Defender', 10),
    ('Diego', 'Carlos', 29, 25000000, 'Defender', 10),
    ('Jules', 'Koundé', 23, 70000000, 'Defender', 10),
    ('Marcos', 'Acuña', 30, 20000000, 'Defender', 10),
    ('Joan', 'Jordan', 27, 35000000, 'Midfielder', 10),
    ('Ivan', 'Rakitić', 34, 5000000, 'Midfielder', 10),
    ('Papu', 'Gómez', 34, 10000000, 'Midfielder', 10),
    ('Lucas', 'Ocampos', 27, 40000000, 'Forward', 10),
    ('Erik', 'Lamela', 31, 25000000, 'Forward', 10),
    ('Youssef', 'En-Nesyri', 24, 45000000, 'Forward', 10);

-- Team: AC Milan #11
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Gianluigi', 'Donnarumma', 24, 60000000, 'Goalkeeper', 11),
    ('Davide', 'Calabria', 26, 30000000, 'Defender', 11),
    ('Fikayo', 'Tomori', 24, 35000000, 'Defender', 11),
    ('Simon', 'Kjær', 32, 20000000, 'Defender', 11),
    ('Theo', 'Hernández', 24, 50000000, 'Defender', 11),
    ('Franck', 'Kessié', 25, 60000000, 'Midfielder', 11),
    ('Ismaël', 'Bennacer', 24, 40000000, 'Midfielder', 11),
    ('Sandro', 'Tonali', 22, 50000000, 'Midfielder', 11),
    ('Ante', 'Rebić', 28, 30000000, 'Forward', 11),
    ('Zlatan', 'Ibrahimović', 40, 5000000, 'Forward', 11),
    ('Rafael', 'Leão', 23, 40000000, 'Forward', 11);

-- Team: Borussia Dortmund #12
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Gregor', 'Kobel', 25, 30000000, 'Goalkeeper', 12),
    ('Thomas', 'Meunier', 31, 10000000, 'Defender', 12),
    ('Manuel', 'Akanji', 26, 40000000, 'Defender', 12),
    ('Mats', 'Hummels', 33, 15000000, 'Defender', 12),
    ('Raphael', 'Guerreiro', 28, 35000000, 'Defender', 12),
    ('Jude', 'Bellingham', 20, 60000000, 'Midfielder', 12),
    ('Axel', 'Witsel', 33, 20000000, 'Midfielder', 12),
    ('Giovanni', 'Reyna', 20, 30000000, 'Midfielder', 12),
    ('Marco', 'Reus', 33, 15000000, 'Forward', 12),
    ('Erling', 'Haaland', 21, 150000000, 'Forward', 12),
    ('Donyell', 'Malen', 23, 40000000, 'Forward', 12);

-- Team: Atletico Madrid #13
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Jan', 'Oblak', 30, 80000000, 'Goalkeeper', 13),
    ('Kieran', 'Trippier', 31, 25000000, 'Defender', 13),
    ('José', 'Giménez', 27, 40000000, 'Defender', 13),
    ('Stefan', 'Savić', 30, 20000000, 'Defender', 13),
    ('Renan', 'Lodi', 23, 30000000, 'Defender', 13),
    ('Jorge', 'Koke', 30, 45000000, 'Midfielder', 13),
    ('Saúl', 'Níguez', 27, 30000000, 'Midfielder', 13),
    ('Héctor', 'Herrera', 32, 10000000, 'Midfielder', 13),
    ('João', 'Félix', 22, 120000000, 'Forward', 13),
    ('Luis', 'Suárez', 35, 5000000, 'Forward', 13),
    ('Ángel', 'Correa', 27, 40000000, 'Forward', 13);

-- Team: Inter Milan #14
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Samir', 'Handanovič', 38, 5000000, 'Goalkeeper', 14),
    ('Alessandro', 'Bastoni', 22, 40000000, 'Defender', 14),
    ('Stefan', 'de Vrij', 30, 40000000, 'Defender', 14),
    ('Milan', 'Škriniar', 27, 50000000, 'Defender', 14),
    ('Achraf', 'Hakimi', 23, 70000000, 'Defender', 14),
    ('Nicolo', 'Barella', 25, 80000000, 'Midfielder', 14),
    ('Marcelo', 'Brozović', 29, 50000000, 'Midfielder', 14),
    ('Hakan', 'Çalhanoğlu', 28, 25000000, 'Midfielder', 14),
    ('Lautaro', 'Martínez', 25, 90000000, 'Forward', 14),
    ('Edin', 'Džeko', 36, 5000000, 'Forward', 14),
    ('Joaquín', 'Correa', 27, 35000000, 'Forward', 14);

-- Team: RB Leipzig #15
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Péter', 'Gulácsi', 32, 20000000, 'Goalkeeper', 15),
    ('Nordi', 'Mukiele', 24, 35000000, 'Defender', 15),
    ('Dayot', 'Upamecano', 23, 60000000, 'Defender', 15),
    ('Lukas', 'Klostermann', 26, 30000000, 'Defender', 15),
    ('Marcel', 'Halstenberg', 30, 20000000, 'Defender', 15),
    ('Dani', 'Olmo', 24, 60000000, 'Midfielder', 15),
    ('Tyler', 'Adams', 23, 30000000, 'Midfielder', 15),
    ('Christopher', 'Nkunku', 24, 40000000, 'Midfielder', 15),
    ('Emil', 'Forsberg', 30, 20000000, 'Forward', 15),
    ('André', 'Silva', 26, 45000000, 'Forward', 15),
    ('Dominik', 'Szoboszlai', 21, 40000000, 'Forward', 15);

-- Team: Arsenal #16
INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id)
VALUES
    ('Bernd', 'Leno', 30, 25000000, 'Goalkeeper', 16),
    ('Héctor', 'Bellerín', 27, 30000000, 'Defender', 16),
    ('Gabriel', 'Magalhães', 24, 40000000, 'Defender', 16),
    ('Ben', 'White', 25, 50000000, 'Defender', 16),
    ('Kieran', 'Tierney', 25, 60000000, 'Defender', 16),
    ('Thomas', 'Partey', 28, 50000000, 'Midfielder', 16),
    ('Granit', 'Xhaka', 30, 25000000, 'Midfielder', 16),
    ('Emile', 'Smith Rowe', 21, 40000000, 'Midfielder', 16),
    ('Bukayo', 'Saka', 20, 80000000, 'Forward', 16),
    ('Pierre-Emerick', 'Aubameyang', 33, 50000000, 'Forward', 16),
    ('Alexandre', 'Lacazette', 30, 30000000, 'Forward', 16);


SELECT * from player;
select * from teams;