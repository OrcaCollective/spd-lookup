CREATE TABLE IF NOT EXISTS portland_officers (
    id                          SERIAL PRIMARY KEY,
    first_name                  VARCHAR(100),
    last_name                   VARCHAR(100),
    gender                      VARCHAR(15),
    rank                        VARCHAR(50),
    employee_id                 VARCHAR(20),
    helmet_id                   VARCHAR(10),
    helmet_id_three_digit       VARCHAR(10),
    salary                      VARCHAR(50),
    
Empl w PPB 3/12/21  ||3||
Empl w PPB 12/28/20  ||3||
Empl w PPB October 2020?  ||3||
Retired/Resigned Since 6/1/2020  ||3||
Retired/Resigned Since 6/1/2020? OR Cert Revoked [Ever]  ||10||
If Yes, Date  ||8||
HIre Year  ||6||
Hire date  ||10||
State Cert Date  ||10||
SupervisoryAdvancedIntermedBasic  ||12||
Employee  (Chest) ID  ||9||
Helmet #  ||4||
3-Dig Helmet #  ||3||
Rank  ||8||
First Name  ||13||
Last Name  ||15||
M/F  ||3||
Badge/DPSST Number  ||7||
Pic on Cops.photo(y/n)  ||3||
RRT(y/n)  ||6||
2016 RRT per 2017 PPB AR  ||3||
2018 Jeff Niiya Email cc RRT  ||3||
2018RRT Specific Training  ||3||
2019RRT Specific Training  ||3||
2020RRT Specific Training  ||3||
2020 Sound Truck Training  ||3||
Has Instructed Course for DPSST 2017+  ||3||
Instructor for Less Lethal/Chem Weapons Courses  ||3||
Cops.Photo Profile Link  ||31||
Has Been Involved in OIS/ Significant UoF Incident  ||3||
Notes  ||498||
Fiscal 2019 Earnings  ||14||
);

COPY portland_officers (last_name,first_name,title,department,salary,date)
FROM '/seed/TPD_Roster_1-24-21.csv' DELIMITER ',' CSV HEADER;

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


CREATE OR REPLACE FUNCTION portland_fuzzy_search_officer_by_name_p(full_name  VARCHAR(100))
    RETURNS SETOF portland_officers AS $$
BEGIN
    RETURN QUERY SELECT *
    FROM portland_officers o
    WHERE LOWER(o.first_name || ' ' || o.last_name) % LOWER(portland_fuzzy_search_officer_by_name_p.full_name)
    ORDER BY SIMILARITY(LOWER(o.first_name || ' ' || o.last_name), LOWER(portland_fuzzy_search_officer_by_name_p.full_name)) DESC;
    RETURN;
END; $$
LANGUAGE 'plpgsql'
SECURITY DEFINER
SET search_path =public, pg_temp;
