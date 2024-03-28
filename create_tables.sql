DROP DATABASE IF EXISTS go;
CREATE DATABASE go;

DROP DATABASE IF EXISTS Users;
DROP DATABASE IF EXISTS Usersdata;

\c go;

CREATE TABLE Users (
    ID SERIAL,
    Username VARCHAR(100) PRIMARY KEY
);

CREATE TABLE Usersdata(
    UserID Int NOT NULL,
    Name VARCHAR(100),
    Surname VARCHAR(100),
    Description VARCHAR(100)
);
