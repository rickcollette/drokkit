#!/bin/bash

# Load environment variables from .env file if it exists
if [ ! -f .env ]; then
    echo ".env file not found. Running drokkgen to generate one..."
    ./drokkgen

    if [ ! -f .env ]; then
        echo "Failed to create .env file. Exiting."
        exit 1
    else
        echo ".env file created successfully."
    fi
else
    echo ".env file found."
fi

# Load environment variables safely
while IFS='=' read -r key value; do
    if [[ ! $key =~ ^# && -n $key ]]; then
        export "$key"="$value"
    fi
done <.env

# Check MySQL status
echo "Checking MySQL connection..."
if mysqladmin ping -h "127.0.0.1" -u "$DB_USERNAME" -p"$DB_PASSWORD" --silent; then
    echo "MySQL is running."
else
    echo "MySQL is not reachable. Please start MySQL and try again."
    exit 1
fi

# Check Redis status
echo "Checking Redis connection..."
redis_host=$(echo "$REDIS_ADDR" | cut -d: -f1)
redis_port=$(echo "$REDIS_ADDR" | cut -d: -f2)

if redis-cli -h "$redis_host" -p "$redis_port" -a "$REDIS_PASSWORD" ping | grep -q "PONG"; then
    echo "Redis is running."
else
    echo "Redis is not reachable. Please start Redis and try again."
    exit 1
fi

# Check NATS status
echo "Checking NATS connection..."
if nats --server "$NATS_URL" ping | grep -q "connected"; then
    echo "NATS is running."
else
    echo "NATS is not reachable. Please start NATS and try again."
    exit 1
fi

# Run drokkit
echo "Starting drokkit..."
./drokkit
