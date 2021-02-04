#!/bin/sh

while getopts udrD flag
do
  case "${flag}" in
    u) up="true";;
    d) down="true";;
    r) restart="true";;
    D) detached="-d";;
    *) error="true";;
  esac
done

if [ $up = "true" ]
then
  docker volume rm spd-lookup_db_data
  docker-compose up --build $detached
fi

if [ $down = "true" ]
then
  docker-compose down
  docker volume rm spd-lookup_db_data
fi

if [ $restart = "true" ]
then
  docker-compose down
  docker volume rm spd-lookup_db_data
  docker-compose up --build $detached
fi

if [ $error = "true" ]
then
  echo "error valid flags are -u, -d, -r"
fi
