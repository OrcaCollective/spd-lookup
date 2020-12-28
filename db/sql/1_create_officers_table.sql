CREATE TABLE IF NOT EXISTS officers (
  id                    SERIAL PRIMARY KEY,
	badge_number          VARCHAR(10),
	first_name            VARCHAR(100),
	middle_name		        VARCHAR(100),
	last_name             VARCHAR(100),
	title                 VARCHAR(100),
	unit                  VARCHAR(50),
  unit_description      VARCHAR(100)
);
