PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE Animals
(
    Name TEXT,
    Sex TEXT,
    BirthDay INTEGER,
    Age INTEGER,
    Species TEXT
);
INSERT INTO Animals VALUES('chiro','female',20080408,14,'cat');
COMMIT;
