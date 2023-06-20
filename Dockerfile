# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables for the PostgreSQL container
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=mydatabase

# Copy SQL scripts to the container
COPY ./pkg/database/init.sql /docker-entrypoint-initdb.d/
EXPOSE 5432
