#!/bin/sh

while getopts udrDf: flag
do
  case "${flag}" in
    u) up="true";;
    d) down="true";;
    r) restart="true";;
    D) detached="-d";;
    f) file="$OPTARG";;
    *) error="true";;
  esac
done

if [ -z "$file" ]
then
  file = "docker-compose.yml"
fi

if [ "$up" = "true" ]
then
  docker volume rm spd-lookup_db_data
  docker-compose -f $file up --build $detached
  chown -R 70:70 db/db_logs
fi

if [ "$down" = "true" ]
then
  docker-compose -f $file down
  docker volume rm spd-lookup_db_data
  rm -rf db/db_logs
fi

if [ "$restart" = "true" ]
then
  docker-compose -f $file down
  docker volume rm spd-lookup_db_data
  rm -rf db/db_logs
  docker-compose -f $file up --build $detached
  chown -R 70:70 db/db_logs
fi

if [ "$error" = "true" ]
then
  echo "error valid flags are -u, -d, -r"
fi
