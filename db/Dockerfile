FROM postgres:alpine

COPY sql/* /docker-entrypoint-initdb.d/
RUN chmod a+r /docker-entrypoint-initdb.d/*

RUN mkdir /pg_logs && chown postgres:postgres /pg_logs

EXPOSE 5432
