FROM postgres:16.1-alpine

RUN chmod -R 775 /var/lib/postgresql/data