CREATE TABLE IF NOT EXISTS olympia_officers (
    id                  SERIAL PRIMARY KEY,
    date                DATE,
    first_name		    VARCHAR(100),
    last_name           VARCHAR(100),
    title               VARCHAR(100),
    unit                VARCHAR(50),
    badge               VARCHAR(10)
);

COPY olympia_officers (date,first_name,last_name,title,unit,badge)
FROM PROGRAM 'curl "https://techblocsea.sfo3.digitaloceanspaces.com/spd-lookup/Olympia-WA-Police-Department_05-01-21.csv"' DELIMITER ',' CSV HEADER;
