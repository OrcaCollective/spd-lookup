FROM postgres:alpine

RUN apt update && apt install fail2ban

COPY sql/* /docker-entrypoint-initdb.d/
RUN chmod a+r /docker-entrypoint-initdb.d/*

EXPOSE 5432
