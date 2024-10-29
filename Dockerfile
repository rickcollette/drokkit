# Use Ubuntu 24.04 as a base image
FROM ubuntu:24.04

# Set environment variables for non-interactive installation
ENV DEBIAN_FRONTEND=noninteractive

# Update package manager and install dependencies
RUN apt-get update && \
    apt-get install -y wget mariadb-server redis-server nginx && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Download and install NATS server
RUN wget https://github.com/nats-io/nats-server/releases/download/v2.10.22/nats-server-v2.10.22-386.deb && \
    dpkg -i nats-server-v2.10.22-386.deb && \
    rm nats-server-v2.10.22-386.deb

# Set the working directory
WORKDIR /app

# Copy application binaries and scripts into the container
COPY dist/drokkit dist/drokkgen dist/runner.sh /app/

# Copy the Nginx configuration
COPY nginx/nginx.conf /etc/nginx/nginx.conf

# Make sure the binaries and shell script have execution permissions
RUN chmod +x /app/drokkit /app/drokkgen /app/runner.sh

# Expose necessary ports (no MariaDB port exposed)
EXPOSE 80 6379 4222 8080

# Copy entrypoint script to initialize services and the application
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Define the command to run on container start
CMD ["/app/entrypoint.sh"]
