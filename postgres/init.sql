CREATE TABLE foods(
    id serial PRIMARY KEY,
    name varchar (50) NOT NULL,
    imagepath varchar(250)
);

CREATE TABLE admins(
    id serial PRIMARY KEY,
    username varchar (50) NOT NULL,
    password varchar(64) NOT NULL
);

CREATE TABLE camps(
    id serial PRIMARY KEY,
    name varchar (50) NOT NULL,
    date bigint,
    processed boolean
);

INSERT INTO foods(name) VALUES('none');

-- DEFAULT USER admin WITH PASSWORD admin. CHANGE THE PASSWORD!!!
INSERT INTO admins(username, password) VALUES('admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918');

INSERT INTO camps(name, date, processed) VALUES('sample', 0, true);