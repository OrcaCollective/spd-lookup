CREATE TABLE IF NOT EXISTS portland_officers (
    id                              SERIAL PRIMARY KEY,
    first_name                      VARCHAR(100),
    last_name                       VARCHAR(100),
    gender                          VARCHAR(15),
    officer_rank                    VARCHAR(50),
    employee_id                     VARCHAR(20),
    helmet_id                       VARCHAR(10),
    helmet_id_three_digit           VARCHAR(10),
    salary                          VARCHAR(50),
    badge                           VARCHAR(20),
    cops_photo_profile_link         VARCHAR(100),
    cops_photo_has_photo            VARCHAR(10),
    employed_3_12_21                VARCHAR(10),
    employed_12_28_20               VARCHAR(10),
    employed_10_01_20               VARCHAR(10),
    retired_6_1_20                  VARCHAR(10),
    retired_or_cert_revoked         VARCHAR(20),
    retired_or_cert_revoked_date    VARCHAR(20),
    hire_year                       VARCHAR(10),
    hire_date                       VARCHAR(20),
    state_cert_date                 VARCHAR(20),
    state_cert_level                VARCHAR(20),
    rrt                             VARCHAR(10),
    rrt_2016                        VARCHAR(10),
    rrt_2018_niiya_email            VARCHAR(10),
    rrt_2018                        VARCHAR(10),
    rrt_2019                        VARCHAR(10),
    rrt_2020                        VARCHAR(10),
    sound_truck_training_2020       VARCHAR(10),
    instructed_for_dpsst            VARCHAR(10),
    instructed_for_less_lethal      VARCHAR(10),
    involved_in_ois_uof             VARCHAR(10),
    notes                           VARCHAR(1000)
);

COPY portland_officers (
    employed_3_12_21,
    employed_12_28_20,
    employed_10_01_20,
    retired_6_1_20,
    retired_or_cert_revoked,
    retired_or_cert_revoked_date,
    hire_year,
    hire_date,
    state_cert_date,
    state_cert_level,
    employee_id,
    helmet_id,
    helmet_id_three_digit,
    officer_rank,
    first_name,
    last_name,
    gender,
    badge,
    cops_photo_has_photo,
    rrt,
    rrt_2016,
    rrt_2018_niiya_email,
    rrt_2018,
    rrt_2019,
    rrt_2020,
    sound_truck_training_2020,
    instructed_for_dpsst,
    instructed_for_less_lethal,
    cops_photo_profile_link,
    involved_in_ois_uof,
    notes,
    salary
)
FROM '/seed/Portland-OR-Police-Bureau_3-20-21.csv' DELIMITER ',' CSV HEADER;

--------------------------------------------------------------------------------
-- Strict search
--------------------------------------------------------------------------------

-- Badge
CREATE OR REPLACE FUNCTION portland_search_officer_by_badge_p(badge VARCHAR(10))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE o.badge = portland_search_officer_by_badge_p.badge;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


-- Employee ID
CREATE OR REPLACE FUNCTION portland_search_officer_by_employee_p(employee_id VARCHAR(20))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE o.employee_id = portland_search_officer_by_employee_p.employee_id;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


-- Helmet ID
CREATE OR REPLACE FUNCTION portland_search_officer_by_helmet_p(helmet_id VARCHAR(10))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE o.helmet_id = portland_search_officer_by_helmet_p.helmet_id;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


-- Helmet ID 3 digit
CREATE OR REPLACE FUNCTION portland_search_officer_by_helmet_three_digit_p(helmet_id_three_digit VARCHAR(10))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE o.helmet_id_three_digit = portland_search_officer_by_helmet_three_digit_p.helmet_id_three_digit;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


-- Name
CREATE OR REPLACE FUNCTION portland_search_officer_by_name_p(
    first_name  VARCHAR(100),
    last_name   VARCHAR(100)
    )
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE LOWER(o.first_name) LIKE LOWER(portland_search_officer_by_name_p.first_name)
    AND LOWER(o.last_name) LIKE LOWER(portland_search_officer_by_name_p.last_name);
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


--------------------------------------------------------------------------------
-- Fuzzy search
--------------------------------------------------------------------------------

-- First name
CREATE OR REPLACE FUNCTION portland_fuzzy_search_officer_by_first_name_p(first_name  VARCHAR(100))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE LOWER(o.first_name) % LOWER(portland_fuzzy_search_officer_by_first_name_p.first_name)
    ORDER BY SIMILARITY(LOWER(o.first_name), LOWER(portland_fuzzy_search_officer_by_first_name_p.first_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;


-- Last name
CREATE OR REPLACE FUNCTION portland_fuzzy_search_officer_by_last_name_p(last_name  VARCHAR(100))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE LOWER(o.last_name) % LOWER(portland_fuzzy_search_officer_by_last_name_p.last_name)
    ORDER BY SIMILARITY(LOWER(o.last_name), LOWER(portland_fuzzy_search_officer_by_last_name_p.last_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;

-- Full Name
CREATE OR REPLACE FUNCTION portland_fuzzy_search_officer_by_name_p(full_name_v  VARCHAR(100))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE LOWER(o.first_name || ' ' || o.last_name) % LOWER(portland_fuzzy_search_officer_by_name_p.full_name_v)
    ORDER BY SIMILARITY(LOWER(o.first_name || ' ' || o.last_name), LOWER(portland_fuzzy_search_officer_by_name_p.full_name_v)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
