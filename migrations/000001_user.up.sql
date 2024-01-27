CREATE TABLE IF NOT  EXISTS user{
    id serial primaty key,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL
}