FROM postgres:alpine

COPY pg_hba.conf /data/
COPY postgresql.conf /data/

COPY sql/* /docker-entrypoint-initdb.d/
RUN chmod a+r /docker-entrypoint-initdb.d/*

EXPOSE 5432
