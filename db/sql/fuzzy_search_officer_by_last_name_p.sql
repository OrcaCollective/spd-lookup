CREATE OR REPLACE FUNCTION fuzzy_search_officer_by_last_name_p(last_name  VARCHAR(100))
    RETURNS SETOF officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM officers o
    WHERE LOWER(o.last_name) % LOWER(fuzzy_search_officer_by_last_name_p.last_name)
    ORDER BY SIMILARITY(LOWER(o.last_name), LOWER(fuzzy_search_officer_by_last_name_p.last_name)) DESC;

    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
