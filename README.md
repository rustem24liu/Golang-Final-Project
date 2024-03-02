# Golang-Final-Project
Final project for discipline Developing in Golang

## Members:
    * Temirgali Rustem 22B030451 
    * Suleimenov Ayan 22B030590 - Team Leader
    * Bagythzhan Zhalgas 22B030317
    * Tolegen Nursultan 22B030453

## Description of project:
Football Teams Management System

Our project revolves around managing football teams participating in tournaments. It encompasses the following key components:

    * List of Teams: 

      This component includes detailed information about several football clubs
      participating in the tournament. Each team entry comprises essential details
      such as the club name, coach, and a comprehensive list of players along with
      their individual data. The implementation involves maintaining a structured 
      list of clubs within the system and providing accessors (getters) to retrieve
      specific information about each club.

    * List of Players:
      This component focuses on maintaining a database of players
      associated with the participating teams. Each player entry contains vital 
      statistics and attributes necessary for team management and performance analysis.

    * Stands:
      This feature provides access to the up-to-date tournament standings
      and results table, allowing users to track team performances throughout the
      tournament. It offers valuable insights into team rankings, points earned, 
      and match outcomes.

    * Schedule: Users can retrieve the detailed schedule of upcoming matches
      through this feature. It provides essential information about match timings, 
      venues, and fixtures, ensuring that users stay informed about the tournament 
      schedule and can plan accordingly.

This project aims to streamline the management of football teams, allowing tournament organizers, teams, and spectators to track match results, schedules, player statistics, and other relevant information efficiently. Each team member is assigned specific tasks based on their skills and interests to ensure the smooth development and functionality of the system.

## Football Management REST API
    * GET:
        GET /matches: Get a list of all matches, add a new match.
        GET /teams: Getting a list of all teams, adding a new team.
        GET /players: Getting a list of all players, adding a new player.
        GET /stands: Getting the tournament results table.
        GET /schedule: Getting the schedule of matches.

    * POST:
        POST /teams: Create a new team, add club name, coach, and players.
        POST /players: Add a new player, include name, position, and stats.
        POST /matches: Create a new match, specify teams, date, and venue.

    * PUT:
        PUT /teams/{team_id}: Update team details, modify club name, coach, or player list.
        PUT /players/{player_id}: Update player information, edit name, position, or statistics.
        PUT /matches/{match_id}: Update match details, revise participating teams, date, or venue.

    * DELETE:
        DELETE /teams/{team_id}: Delete a team from the tournament.
        DELETE /players/{player_id}: Remove a player from the system.
        DELETE /matches/{match_id}: Cancel a match and remove it from the tournament schedule.

## DB Structure

    Table coach {
    id integer [primary key]
    firstname varchar
    secondname varchar
    experience_year integer
    }
    
    Table players {
    id integer [primary key]
    firstname varchar
    secondname varchar
    age integer
    cost decimal
    position varchar
    team_id integer [unique, ref: > teams.id]
    }
    
    Table teams {
    id integer [primary key]
    name varchar
    coach_id integer [unique, ref: > coach.id]
    }
    
    Ref: players.team_id - teams.id (one-to-one)
    Ref: teams.coach_id - coach.id (one-to-one)
