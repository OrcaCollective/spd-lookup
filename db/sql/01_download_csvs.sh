#!/bin/bash
curl "$ROSTER_SOURCE/seattle.csv" -o /tmp/seattle.csv
curl "$ROSTER_SOURCE/tacoma.csv" -o /tmp/tacoma.csv
curl "$ROSTER_SOURCE/portland.csv" -o /tmp/portland.csv
curl "$ROSTER_SOURCE/auburn.csv" -o /tmp/auburn.csv
curl "$ROSTER_SOURCE/lakewood.csv" -o /tmp/lakewood.csv
curl "$ROSTER_SOURCE/olympia.csv" -o /tmp/olympia.csv
curl "$ROSTER_SOURCE/bellevue.csv" -o /tmp/bellevue.csv
curl "$ROSTER_SOURCE/renton.csv" -o /tmp/renton.csv
curl "$ROSTER_SOURCE/thurston_co.csv" -o /tmp/thurston_co.csv
curl "$ROSTER_SOURCE/port_of_seattle.csv" -o /tmp/port_of_seattle.csv