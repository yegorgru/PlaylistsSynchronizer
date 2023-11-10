CREATE TABLE users
(
    ID            serial PRIMARY KEY,
    username          varchar(255) not null unique,
    email      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE groups
(
    ID serial PRIMARY KEY,
    name varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE roles
(
    ID serial PRIMARY KEY,
    name varchar(255) not null
);

CREATE TABLE user_group
(
    ID serial PRIMARY KEY,
    userID serial REFERENCES users(ID) ON DELETE CASCADE not null,
    groupID serial REFERENCES groups(ID) ON DELETE CASCADE not null,
    roleID serial REFERENCES roles(ID) ON DELETE CASCADE not null
);

CREATE TABLE playlists
(
    ID serial PRIMARY KEY,
    name varchar(255) not null,
    groupID serial REFERENCES groups(ID) ON DELETE CASCADE not null
);

