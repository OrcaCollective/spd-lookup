CREATE OR REPLACE FUNCTION search_officer_by_name_p(
    first_name  VARCHAR(100),
    last_name VARCHAR(100)
    )
    RETURNS SETOF officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM officers o
    WHERE lower(o.first_name) like lower(search_officer_by_name_p.first_name)
    AND lower(o.last_name) like lower(search_officer_by_name_p.last_name);

    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
