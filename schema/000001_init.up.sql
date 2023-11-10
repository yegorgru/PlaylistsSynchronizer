CREATE TABLE users
(
    id            serial PRIMARY KEY,
    username          varchar(255) not null unique,
    email      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE groups
(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE roles
(
    id serial PRIMARY KEY,
    name varchar(255) not null
);

CREATE TABLE user_group
(
    id serial PRIMARY KEY,
    userID serial REFERENCES users(id) ON DELETE CASCADE not null,
    groupID serial REFERENCES groups(id) ON DELETE CASCADE not null,
    roleID serial REFERENCES roles(id) ON DELETE CASCADE not null
);

CREATE TABLE playlists
(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    groupID serial REFERENCES groups(id) ON DELETE CASCADE not null
);

