COPY officers (badge_number,title,unit,unit_description,last_name,first_name,middle_name)
FROM '/seed/SPD_roster_11-15-20.csv' DELIMITER ',' CSV HEADER;

COPY tacoma_officers (last_name,first_name,title,department,salary)
FROM '/seed/tacoma_roster_012421.csv' DELIMITER ',' CSV HEADER;
