# FROM postgres:16.0

# COPY logdb.sql /docker-entrypoint-initdb.d/

# RUN ["apt-get", "update"]
# RUN ["apt-get", "install", "-y", "vim"]


FROM postgres:16.0 as tmpImg
COPY logdb.sql /docker-entrypoint-initdb.d/

RUN ["sed", "-i", "s/exec \"$@\"/echo \"skipping...\"/", "/usr/local/bin/docker-entrypoint.sh"]

ENV POSTGRES_USER=ziyan
ENV POSTGRES_PASSWORD=postgres
ENV PGDATA=/data

RUN ["/usr/local/bin/docker-entrypoint.sh", "postgres"]

FROM postgres:16.0
COPY --from=tmpImg /data $PGDATA
