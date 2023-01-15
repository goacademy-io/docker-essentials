FROM postgres:14.5
LABEL maintainer="info@goacademy.io"
COPY world.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_PASSWORD=root
EXPOSE 5432