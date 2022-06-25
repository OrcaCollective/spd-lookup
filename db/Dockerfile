FROM postgres:alpine

RUN apk --update add curl --no-cache

COPY ./db/*.conf /data/

COPY ./db/sql/* /docker-entrypoint-initdb.d/

EXPOSE 5432
