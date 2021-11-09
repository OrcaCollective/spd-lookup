CREATE TABLE IF NOT EXISTS bellevue_officers (
    id                  SERIAL PRIMARY KEY,
    title   		    VARCHAR(100),
    last_name           VARCHAR(100),
    first_name          VARCHAR(100),
    unit                VARCHAR(50),
    notes               VARCHAR(100),
    badge               VARCHAR(50)
);

COPY bellevue_officers (first_name,last_name,title,badge,unit,notes)
FROM '/seed/Bellevue-WA-Police-Department_05-01-21.csv' DELIMITER ',' CSV HEADER;
