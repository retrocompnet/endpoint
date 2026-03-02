CREATE TABLE Settings (
       key VARCHAR(255),
       val VARCHAR(255)
);

UPDATE Migration SET db_version=0001 WHERE db_version=0000;
