CREATE TABLE IF NOT EXISTS port_of_seattle_officers (
    id            SERIAL PRIMARY KEY,
    badge_number  VARCHAR(50),
    name          VARCHAR(100),
    rank          VARCHAR(50),
    hire_date     INTEGER,
    unit          VARCHAR(100)
);

COPY port_of_seattle_officers (badge_number,name,hire_date,rank,unit)
FROM PROGRAM 'curl "https://techblocsea.sfo3.digitaloceanspaces.com/spd-lookup/PortOfSeattle-WA-Police-Department_05-01-21.csv"' DELIMITER ',' CSV HEADER;
