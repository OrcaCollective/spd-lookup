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
FROM '/tmp/olympia.csv' DELIMITER ',' CSV HEADER;
