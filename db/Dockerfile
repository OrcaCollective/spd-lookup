FROM postgres:alpine

COPY *.conf /data/

COPY sql/* /docker-entrypoint-initdb.d/
RUN chmod a+r /docker-entrypoint-initdb.d/*

COPY seed/* /seed/
RUN chmod a+r /seed/*

EXPOSE 5432
