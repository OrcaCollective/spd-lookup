CREATE TABLE IF NOT EXISTS tacoma_officers (
    id              SERIAL PRIMARY KEY,
    date            DATE,
    first_name      VARCHAR(100),
    last_name       VARCHAR(100),
    title           VARCHAR(100),
    department      VARCHAR(100),
    salary          VARCHAR(50)
);

COPY tacoma_officers (last_name,first_name,title,department,salary,date)
FROM PROGRAM 'curl "https://techblocsea.sfo3.digitaloceanspaces.com/spd-lookup/Tacoma-WA-Police-Department_1-24-21.csv"' DELIMITER ',' CSV HEADER;

CREATE OR REPLACE FUNCTION tacoma_search_officer_by_name_p(
    first_name  VARCHAR(100),
    last_name   VARCHAR(100)
    )
    RETURNS SETOF tacoma_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM tacoma_officers o
    WHERE LOWER(o.first_name) LIKE LOWER(tacoma_search_officer_by_name_p.first_name)
    AND LOWER(o.last_name) LIKE LOWER(tacoma_search_officer_by_name_p.last_name);
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION tacoma_fuzzy_search_officer_by_first_name_p(first_name  VARCHAR(100))
    RETURNS SETOF tacoma_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM tacoma_officers o
    WHERE LOWER(o.first_name) % LOWER(tacoma_fuzzy_search_officer_by_first_name_p.first_name)
    ORDER BY SIMILARITY(LOWER(o.first_name), LOWER(tacoma_fuzzy_search_officer_by_first_name_p.first_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION tacoma_fuzzy_search_officer_by_last_name_p(last_name  VARCHAR(100))
    RETURNS SETOF tacoma_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM tacoma_officers o
    WHERE LOWER(o.last_name) % LOWER(tacoma_fuzzy_search_officer_by_last_name_p.last_name)
    ORDER BY SIMILARITY(LOWER(o.last_name), LOWER(tacoma_fuzzy_search_officer_by_last_name_p.last_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION tacoma_fuzzy_search_officer_by_name_p(full_name  VARCHAR(100))
    RETURNS SETOF tacoma_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM tacoma_officers o
    WHERE LOWER(o.first_name || ' ' || o.last_name) % LOWER(tacoma_fuzzy_search_officer_by_name_p.full_name)
    ORDER BY SIMILARITY(LOWER(o.first_name || ' ' || o.last_name), LOWER(tacoma_fuzzy_search_officer_by_name_p.full_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
