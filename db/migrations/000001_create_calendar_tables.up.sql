CREATE TYPE subscribe_types AS ENUM (
  'creator',
  'admin',
  'viewer'
);


CREATE TABLE users
(
    id serial NOT NULL UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    is_active bool NOT NULL DEFAULT false
);

CREATE TABLE sessions
(
    id serial NOT NULL UNIQUE,
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    refresh_token varchar(255) NOT NULL UNIQUE,
    fingerprint varchar(255) NOT NULL UNIQUE,
    issued_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at timestamp NOT NULL
);

CREATE TABLE calendars
(
    id serial NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    description varchar NOT NULL
);

CREATE TABLE subscribes
(
    id serial NOT NULL UNIQUE,
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    calendar_id int REFERENCES calendars(id) ON DELETE CASCADE NOT NULL,
    type subscribe_types NOT NULL DEFAULT 'viewer'
);


CREATE TABLE events
(
    id serial NOT NULL UNIQUE,
    calendar_id int REFERENCES calendars(id) ON DELETE CASCADE NOT NULL,
    title varchar(255) NOT NULL,
    description varchar NOT NULL,
    time timestamp NOT NULL
);