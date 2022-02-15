CREATE TABLE IF NOT EXISTS lakewood_officers (
    id                  SERIAL PRIMARY KEY,
    date                DATE,
    title   		    VARCHAR(100),
    last_name           VARCHAR(100),
    first_name          VARCHAR(100),
    unit                VARCHAR(50),
    unit_description    VARCHAR(100)
);

COPY lakewood_officers (date,title,last_name,first_name,unit,unit_description)
FROM PROGRAM 'curl "https://techblocsea.sfo3.digitaloceanspaces.com/spd-lookup/Lakewood-WA-Police-Department_05-01-21.csv"' DELIMITER ',' CSV HEADER;
