FROM postgres:9.4

COPY sql/* /docker-entrypoint-initdb.d/

EXPOSE 5432

#psql -h localhost -p 5432 -U postgres