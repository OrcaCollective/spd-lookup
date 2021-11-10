CREATE TABLE IF NOT EXISTS thurston_officers (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(50),
    first_name  VARCHAR(100),
    last_name   VARCHAR(100),
    call_sign   VARCHAR(50),
    call_sign_2 VARCHAR(50)
);

COPY thurston_officers (title,last_name,first_name,call_sign,call_sign_2)
FROM '/seed/ThurstonCounty-WA-Sheriffs-Office_05-01-21.csv' DELIMITER ',' CSV HEADER;
