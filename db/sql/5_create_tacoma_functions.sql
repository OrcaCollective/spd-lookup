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
