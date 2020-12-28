COPY officers (badge_number,title,unit,unit_description,last_name,first_name,middle_name)
FROM '/seed/SPD_roster_11-15-20.csv' DELIMITER ',' CSV HEADER;
