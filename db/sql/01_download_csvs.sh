#!/bin/bash
curl "$ROSTER_SOURCE/Seattle-WA-Police-Department_Historical.csv" -o /tmp/seattle.csv
curl "$ROSTER_SOURCE/Tacoma-WA-Police-Department_1-24-21.csv" -o /tmp/tacoma.csv
curl "$ROSTER_SOURCE/Portland-OR-Police-Bureau_3-20-21.csv" -o /tmp/portland.csv
curl "$ROSTER_SOURCE/Auburn-WA-Police-Department_06-07-21.csv" -o /tmp/auburn.csv
curl "$ROSTER_SOURCE/Lakewood-WA-Police-Department_05-01-21.csv" -o /tmp/lakewood.csv
curl "$ROSTER_SOURCE/Olympia-WA-Police-Department_05-01-21.csv" -o /tmp/olympia.csv
curl "$ROSTER_SOURCE/Bellevue-WA-Police-Department_05-01-21.csv" -o /tmp/bellevue.csv
curl "$ROSTER_SOURCE/Renton-WA-Police-Department_05-01-21.csv" -o /tmp/renton.csv
curl "$ROSTER_SOURCE/ThurstonCounty-WA-Sheriffs-Office_05-01-21.csv" -o /tmp/thurston_co.csv
curl "$ROSTER_SOURCE/PortOfSeattle-WA-Police-Department_05-01-21.csv" -o /tmp/port_of_seattle.csv