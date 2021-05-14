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
FROM '/seed/SPD_Rosters_Historical.csv' DELIMITER ',' CSV HEADER;

CREATE OR REPLACE FUNCTION seattle_get_officer_by_badge_p(badge VARCHAR(10))
    RETURNS SETOF seattle_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM seattle_officers o
    WHERE o.badge = seattle_get_officer_by_badge_p.badge;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION seattle_search_officer_by_name_p(
    first_name  VARCHAR(100),
    last_name VARCHAR(100)
    )
    RETURNS SETOF seattle_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM seattle_officers o
    WHERE LOWER(o.first_name) LIKE LOWER(seattle_search_officer_by_name_p.first_name)
    AND LOWER(o.last_name) LIKE LOWER(seattle_search_officer_by_name_p.last_name)
    ORDER BY o.date DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION seattle_fuzzy_search_officer_by_first_name_p(first_name  VARCHAR(100))
    RETURNS SETOF seattle_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM seattle_officers o
    WHERE LOWER(o.first_name) % LOWER(seattle_fuzzy_search_officer_by_first_name_p.first_name)
    ORDER BY
        o.date DESC,
        SIMILARITY(LOWER(o.first_name), LOWER(seattle_fuzzy_search_officer_by_first_name_p.first_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION seattle_fuzzy_search_officer_by_last_name_p(last_name  VARCHAR(100))
    RETURNS SETOF seattle_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM seattle_officers o
    WHERE LOWER(o.last_name) % LOWER(seattle_fuzzy_search_officer_by_last_name_p.last_name)
    ORDER BY
        o.date DESC, 
        SIMILARITY(LOWER(o.last_name), LOWER(seattle_fuzzy_search_officer_by_last_name_p.last_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


CREATE OR REPLACE FUNCTION seattle_fuzzy_search_officer_by_name_p(full_name_v  VARCHAR(100))
    RETURNS SETOF seattle_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM seattle_officers o
    WHERE LOWER(o.first_name || ' ' || o.last_name) % LOWER(seattle_fuzzy_search_officer_by_name_p.full_name_v)
    ORDER BY
        o.date DESC, 
        SIMILARITY(LOWER(o.first_name || ' ' || o.last_name), LOWER(seattle_fuzzy_search_officer_by_name_p.full_name_v)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
