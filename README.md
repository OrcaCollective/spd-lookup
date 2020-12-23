# SPD Lookup
Project that allows for searching for SPD police officers by badge, first_name, or last_name. Currently, API and database are hosted on heroku; API is avialable at [https://spd-lookup.herokuapp.com](https://spd-lookup.herokuapp.com)

## Resources
- `/officer` - expects `badge`, `first_name` and/or `last_name` to be provided as query parameters. An array of officers will be returned
  - if `badge` is provided, will look up officer in database by badge
  - else if either `first_name` or `last_name`, a name search will be performed on the database. Due to URL encoding, `*` will be treated as a wildcard
- `/officer/search` - expects `first_name` and/or `last_name` to be provided as query parameters. Invokes a fuzzy match based on name, array of officers returned are in descending match score

## API
### Running Locally
To start the server, the following environment variables need to be provided
1. `DB_HOST`: database host name
1. `DB_NAME`: database name
1. `DB_USERNAME`: user to connect to database
1. `DB_PASSWORD`: password to connect to database
1. `PORT`: port to listen for
sample usage:
```
cd api
PORT=5000 \
DB_USERNAME=your_username \
DB_PASSWORD=your_password \
DB_NAME=your_db_name \
DB_HOST=your_db_host \
go run *.go
```

## Database
### DB Seed Script

Quick script to read from a CSV file and load data into an officers table in a postgres database. Assumes seed csv file contains current SPD roster, as the script first clears the table before inserting values.

### Usage
Script expects the following environment variables
1. `SEED_FILE`: relative path to seed csv file
1. `DB_HOST`: database host name
1. `DB_NAME`: database name
1. `DB_USERNAME`: user to connect to database
1. `DB_PASSWORD`: password to connect to database

sample usage:
```
cd db
SEED_FILE=csv/SPD_roster_11-15-20.csv \
DB_USERNAME=your_username \
DB_PASSWORD=your_password \
DB_NAME=your_db_name \
DB_HOST=your_db_host \
go run main.go
```
