FROM postgres:alpine
RUN apt-get update
COPY docker-entrypoint.sh /usr/local/bin/
COPY create_user.sh   /docker-entrypoint-initdb.d/10-create_user.sh
COPY init.sql  /docker-entrypoint-initdb.d/20-create_user.sql
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

