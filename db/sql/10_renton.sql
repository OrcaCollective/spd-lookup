CREATE TABLE IF NOT EXISTS renton_officers (
    id                  SERIAL PRIMARY KEY,
    last_name           VARCHAR(100),
    first_name          VARCHAR(100),
    middle_name         VARCHAR(100),
    rank                VARCHAR(100),
    department          VARCHAR(100),
    division            VARCHAR(100),
    shift               VARCHAR(100),
    additional_info     VARCHAR(100),
    badge_number        VARCHAR(100)
);

COPY renton_officers (last_name,first_name,middle_name,rank,department,division,shift,additional_info,badge_number)
FROM '/tmp/renton.csv' DELIMITER ',' CSV HEADER;
