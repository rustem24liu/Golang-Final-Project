-- Filename: 001_CreateTables.sql

CREATE TABLE IF NOT EXISTS Teams (
                                     team_id SERIAL PRIMARY KEY,
                                     team_name VARCHAR(100)
    );

CREATE TABLE IF NOT EXISTS Player (
                                      player_id SERIAL PRIMARY KEY,
                                      first_name VARCHAR(50),
    last_name VARCHAR(50),
    player_age INTEGER,
    player_cost DECIMAL,
    player_pos VARCHAR(50),
    team_id INTEGER,
    FOREIGN KEY (team_id) REFERENCES Teams(team_id)
    );

CREATE TABLE IF NOT EXISTS Coach (
                                     coach_id SERIAL PRIMARY KEY,
                                     first_name VARCHAR(50),
    last_name VARCHAR(50),
    exp_year INTEGER,
    team_id INTEGER,
    FOREIGN KEY (team_id) REFERENCES Teams(team_id)
    );

CREATE TABLE IF NOT EXISTS users (
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
