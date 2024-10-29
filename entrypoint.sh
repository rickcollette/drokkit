#!/bin/bash

# Start MariaDB, Redis, NATS server, and Nginx in the background
service mysql start
service redis-server start
nats-server &
service nginx start

# Check if .env file exists; if not, generate it
if [ ! -f .env ]; then
    echo ".env file not found. Running drokkgen to generate one..."
    ./drokkgen
fi

# Run the main application setup and server script
exec ./runner.sh
