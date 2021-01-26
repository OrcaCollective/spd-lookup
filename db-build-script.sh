#!/bin/sh

docker-compose down
docker volume rm spd-lookup_db_data
docker-compose up --build