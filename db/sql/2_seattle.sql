CREATE TABLE IF NOT EXISTS seattle_officers (
    id                  SERIAL PRIMARY KEY,
    date                DATE,
    full_name		    VARCHAR(100),
    badge               VARCHAR(10),
    first_name          VARCHAR(100),
    middle_name         VARCHAR(100),
    last_name           VARCHAR(100),
    title               VARCHAR(100),
    unit                VARCHAR(50),
    unit_description    VARCHAR(100)
);

COPY seattle_officers (badge,full_name,title,unit,unit_description,first_name,middle_name,last_name,date)
FROM '/seed/Seattle-WA-Police-Department_Historical.csv' DELIMITER ',' CSV HEADER;
