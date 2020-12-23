CREATE OR REPLACE FUNCTION fuzzy_search_officer_by_first_name_p(first_name  VARCHAR(100))
    RETURNS SETOF officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM officers o
    WHERE LOWER(o.first_name) % LOWER(fuzzy_search_officer_by_first_name_p.first_name)
    ORDER BY SIMILARITY(LOWER(o.first_name), LOWER(fuzzy_search_officer_by_first_name_p.first_name)) DESC;

    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
