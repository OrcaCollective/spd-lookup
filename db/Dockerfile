FROM postgres:alpine

COPY *.conf /data/

COPY sql/* /docker-entrypoint-initdb.d/

COPY seed/* /seed/

EXPOSE 5432
