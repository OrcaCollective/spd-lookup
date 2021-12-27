FROM postgres:alpine

COPY ./db/*.conf /data/

COPY ./db/sql/* /docker-entrypoint-initdb.d/

COPY ./db/seed/* /seed/

EXPOSE 5432
