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

CREATE OR REPLACE FUNCTION fuzzy_search_officer_p(badge_number  VARCHAR(10), last_name  VARCHAR(100), first_name  VARCHAR(100))
    RETURNS SETOF officers AS $$
BEGIN
    RETURN QUERY SELECT *
                 FROM officers o
                 WHERE LOWER(o.badge_number) % LOWER(fuzzy_search_officer_p.badge_number)
                 OR LOWER(o.last_name) % LOWER(fuzzy_search_officer_p.last_name)
                 OR LOWER(o.first_name) % LOWER(fuzzy_search_officer_p.first_name)
                 ORDER BY SIMILARITY(LOWER(o.badge_number), LOWER(fuzzy_search_officer_p.badge_number)) DESC,
                          (.5 * SIMILARITY(LOWER(o.last_name), LOWER(fuzzy_search_officer_p.last_name))) DESC,
                          SIMILARITY(LOWER(o.first_name), LOWER(fuzzy_search_officer_p.first_name)) DESC;

    RETURN;
END; $$
    LANGUAGE 'plpgsql'
    SECURITY DEFINER
    SET search_path =public, pg_temp;
