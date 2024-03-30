# Go Testing Proof of Concept (POC)

## Overview

This repository serves as a proof of concept for testing the implementation of a web application using the Go programming language.

## Prerequisites

Ensure that Go is installed on your system. You can download and install it from the [official Go website](https://golang.org/).
Also, make sure you have Docker and Docker Compose installed.

## Usage

### Running the Server

Navigate to the root directory of the project in your terminal:

```bash
cd {{ root_path }}
```

Starting the PostgreSQL server:

```bash
docker-compose up -d postgres
```

Then, execute the following command to start the server:

```bash
go run .
```

## Database Migration

For the database migration, we use golang-migrate/migrate. For more information, refer to the [GitHub repository](https://github.com/golang-migrate/migrate) or the [tutourial](https://www.freecodecamp.org/news/database-migration-golang-migrate/)

### Install Migration CLI

To install the migrate CLI tool using Scoop on Windows, follow these steps:

```bash
scoop install migrate
```

To install the migrate CLI tool on Mac, follow these steps:

```bash
brew install golang-migrate
```

### Create a New Migration

When adding new features or modifying the database schema, it's essential to create a migration to manage these changes effectively. To create a new migration, use the following command:

```bash
migrate create -ext sql -dir database/migration/ -seq {{ migration_name }}
```

### Applied the Migration

After creating a migration, it needs to be applied to the database to enact the changes. Use the following command to apply the migration:

```bash
migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
```

## API Documentation

We leverage [swaggo/swag](https://github.com/swaggo/swag) for our API documentation.  The Swagger documentation is accessible at the following path: `http://localhost:8080/docs/index.html`.

### Installation

You can install the tool using the following command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Updating API Description

To update the API description, execute the following command at the root directory:

```bash
swag init
```

This command will generate or update the API documentation based on the comments in your Go code. Make sure to annotate your code appropriately for swaggo to parse and generate the documentation correctly.

## License

This project is licensed under the terms of the MIT license. See the LICENSE file for details.
