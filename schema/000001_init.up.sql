CREATE TABLE League(
                       id serial primary key,
                       league_name varchar(250)
);


CREATE TABLE Teams (
                       team_id SERIAL PRIMARY KEY,
                       team_name VARCHAR(100),
                       league_id INT,
                       foreign key (league_id) references league(id)
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

CREATE TABLE Stadiums (
                          id SERIAL PRIMARY KEY,
                          stadium_name VARCHAR(100),
                          capacity INT,
                          team_id INT,
                          FOREIGN KEY (team_id) REFERENCES teams(team_id)
);