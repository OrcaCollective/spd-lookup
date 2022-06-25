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
FROM '/tmp/lakewood.csv' DELIMITER ',' CSV HEADER;
