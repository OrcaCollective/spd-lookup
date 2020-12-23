CREATE OR REPLACE FUNCTION search_officer_by_name_p(
    first_name  VARCHAR(100),
    last_name VARCHAR(100)
    )
    RETURNS SETOF officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM officers o
    WHERE LOWER(o.first_name) LIKE LOWER(search_officer_by_name_p.first_name)
    AND LOWER(o.last_name) LIKE LOWER(search_officer_by_name_p.last_name);

    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
