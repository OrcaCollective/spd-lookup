## DB Seed Script

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
SEED_FILE=csv/SPD_roster_11-15-20.csv \
DB_USERNAME=your_username \
DB_PASSWORD=your_password \
DB_NAME=your_db_name \
DB_HOST=your_db_host \
go run main.go
```
