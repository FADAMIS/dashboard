CREATE TABLE foods(
    id serial PRIMARY KEY,
    name varchar (50) NOT NULL,
    imagepath varchar(250)
);

INSERT INTO foods(name) VALUES('none');