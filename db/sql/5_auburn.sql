CREATE TABLE IF NOT EXISTS auburn_officers (
    id                  SERIAL PRIMARY KEY,
    date                DATE,
    last_name		    VARCHAR(100),
    first_name          VARCHAR(100),
    badge               VARCHAR(10),
    title               VARCHAR(100)
);

COPY auburn_officers (date,last_name,first_name,badge,title)
FROM PROGRAM 'curl "https://techblocsea.sfo3.digitaloceanspaces.com/spd-lookup/Auburn-WA-Police-Department_06-07-21.csv"' DELIMITER ',' CSV HEADER;
