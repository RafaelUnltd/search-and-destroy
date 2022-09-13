-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE MATCH_STATUS AS ENUM ('making', 'in_progress', 'finished');
CREATE TYPE TEAM AS ENUM ('ct', 'tr');

CREATE TABLE IF NOT EXISTS players (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS matches (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    status MATCH_STATUS NOT NULL default 'making',
    winner_team TEAM default NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS players_matches (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_id UUID NOT NULL,
    match_id UUID NOT NULL,
    team TEAM NOT NULL,
    kills integer NOT NULL DEFAULT 0,
    deaths integer NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_player FOREIGN KEY(player_id) REFERENCES players(id),
    CONSTRAINT fk_match FOREIGN KEY(match_id) REFERENCES matches(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players_matches;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS matches;
DROP TYPE IF EXISTS MATCH_STATUS;
DROP TYPE IF EXISTS TEAM;
-- +goose StatementEnd
