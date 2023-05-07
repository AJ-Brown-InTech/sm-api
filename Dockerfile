# Use a small Linux distribution as the base image
FROM alpine:3.14

# Install SQLite and its command-line tool
RUN apk add --no-cache sqlite sqlite-dev

# Copy the SQLite database file into the container
COPY database.db /app/database.db

# Set the working directory to /app
WORKDIR /app

# Expose port 8080 (or whatever port your application uses)
EXPOSE 8080

# Start the SQLite server and listen on port 8080
CMD ["sqlite3", "database.db"]
