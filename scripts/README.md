# Preparation scripts

This folder contains two scripts which help clean and ingest rosters provided by SPD. The rosters typically come in the form of XLSX files that must first be converted to CSVs manually. Once that step is done, the following scripts can be run.

## Prepare SPD roster

This script will clean a CSV which was generated from the XLSX file. It also adds a date field to all records in the CSV. This is a prerequisite to adding the roster to the historical aggregate.

**File**: `prep-spd-roster.py`
**Help**: 
```
Usage: prep-spd-roster.py [OPTIONS]

Options:
  -d, --date TEXT     The date of the roster in the format YYYY-MM-DD.
                      [required]
  -i, --in-csv TEXT   Path to the source file. Can be relative to current
                      directory.  [required]
  -o, --out-csv TEXT  Path to save the resulting CSV in. Can be relative to
                      current directory.  [required]
  --help              Show this message and exit.
```
**Example**: `./prep-spd-roster.py -d 2021-11-10 -i ~/data/roster-2021-11-10.csv -o ~/data/roster-2021-11-10_processed.csv`

## Add to current historical roster

This script will take a CSV that has been prepared by `prep-spd.roster.py` and add it to a historical roster (defaults to `../db/seed/Seattle-WA-Police-Department_Historical.csv`).

**File**: `add-to-historical-roster.py`
**Help**:
```
Usage: add-to-historical-roster.py [OPTIONS] [INPUT_CSVS]...

Options:
  -h, --historical-csv TEXT  Path to the historical CSV. This file will be
                             replaced with the updated version
  --help                     Show this message and exit.
```
**Example**: `./add-to-historical-roster.py ~/data/roster-2021-11-10_processed.csv`
