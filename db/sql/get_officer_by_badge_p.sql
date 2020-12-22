CREATE OR REPLACE FUNCTION get_officer_by_badge_p(badge_number VARCHAR(10))
    RETURNS SETOF officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM officers o
    WHERE o.badge_number = get_officer_by_badge_p.badge_number;

    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
