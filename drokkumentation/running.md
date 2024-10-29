# Drokkit Game Server Installation Guide

This guide walks you through the installation, configuration, and verification process for running the Drokkit game server. Follow these steps to install and set up MariaDB, Redis, and NATS, and to configure your environment.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installing MariaDB](#installing-mariadb)
- [Installing Redis](#installing-redis)
- [Installing NATS](#installing-nats)
- [Environment Configuration](#environment-configuration)
- [Running the Application](#running-the-application)

## Prerequisites

Ensure you have the following installed on your system:

- **Go (Golang)** - [Install Go](https://golang.org/doc/install)
- **Git** - [Install Git](https://git-scm.com/downloads)

## Environment Configuration

The `.env` file stores configuration variables needed by the application, such as database connection details and secret keys.  You'll want to run this first, since it contains all the relevant information for your databases.

### Create or Generate the .env File

Use the `drokkgen` tool to generate a new `.env` file or manually create one.
**NOTE:** `drokkgen` uses uuid and base64 to create the password for your databases.

Here's a sample configuration:

```env
# Main Relational DB
DB_DRIVER=mysql
DB_SOURCE=drokkit:your_password@tcp(localhost:3306)/megacity?charset=utf8mb4&parseTime=True&loc=Local
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100
DB_CONN_MAX_LIFETIME=30m

# Redis DB for NATS and Leaderboards
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=your_redis_password

# JWT secret key
JWT_SECRET_KEY=your_jwt_secret_key

# NATS configuration
NATS_URL=nats://localhost:4222
```

## Installing MariaDB

### 1. Install MariaDB

```bash
sudo apt update
sudo apt install mariadb-server -y
```

### 2. Secure the Installation

Run the MariaDB secure installation script to set a root password and remove insecure defaults:

```bash
sudo mysql_secure_installation
```

### 3. Start and Enable MariaDB Service

```bash
sudo systemctl start mariadb
sudo systemctl enable mariadb
```

### 4. Create the Database and User

Connect to MariaDB:

```bash
sudo mysql -u root -p
```

Use the information found in your `.env` file to fill in the password information.

Run the following SQL commands:

```sql
CREATE DATABASE megacity CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'drokkit'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON megacity.* TO 'drokkit'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

## Installing Redis

### 1. Install Redis

```bash
sudo apt update
sudo apt install redis-server -y
```

### 2. Configure Redis for Persistence (Optional)

Edit the Redis configuration file:

```bash
sudo nano /etc/redis/redis.conf
```

Uncomment and adjust the following line for persistence:

```bash
appendonly yes
```

### 3. Start and Enable Redis Service

```bash
sudo systemctl start redis-server
sudo systemctl enable redis-server
```

### 4. Verify Redis Installation

```bash
redis-cli ping
```

You should see `PONG` if Redis is running correctly.

## Installing NATS

### 1. Download NATS Server

```bash
wget https://github.com/nats-io/nats-server/releases/download/v2.10.22/nats-server-v2.10.22-386.deb
sudo dpkg -i nats-server-v2.10.22-386.deb
```

### 2. Run NATS Server

```bash
sudo systemctl start nats
```

### 3. Verify NATS Installation

```bash
nats --server nats://localhost:4222 ping
```

## Running the Application

### 1. Two ways to do this

1. Follow the instructions above.
2. Run the `drokkit` server

```bash
./drokkit &
```

**Or:**

Ensure the script verifies or generates the `.env` file, checks dependencies, and starts the server:

```bash
./runner.sh
```

### 2. Manual Commands for Starting Services and Verifications

Check MariaDB Connection:

```bash
mysqladmin ping -h 127.0.0.1 -u drokkit -pYourPassword --silent
```

Check Redis Connection:

```bash
redis-cli -h $(echo $REDIS_ADDR | cut -d: -f1) -p $(echo $REDIS_ADDR | cut -d: -f2) -a $REDIS_PASSWORD ping
```

Check NATS Connection:

```bash
nats --server $NATS_URL ping
```

### 3. Start the Application

```bash
./drokkit
```

Your Drokkit game server should now be running and ready for connections!
