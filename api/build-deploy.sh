#!/bin/sh

heroku container:push web -a spd-lookup
heroku container:release web -a spd-lookup -v
