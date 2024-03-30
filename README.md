# go-testing-poc

## Overview

This repository serves as a proof of concept for testing the implementation of a web application using the Golang framework.

## Prerequisites

Ensure you have Go installed on your system. You can download and install it from the [official Go website](https://golang.org/).
Also, make sure you have Docker and Docker Compose installed.

## Usage

### Running the Server

Navigate to the root directory of the project in your terminal:

```bash
cd {{ root_path }}
```

(Optional) Starting the PostgreSQL server:

```bash
docker-compose up -d postgres
```

Then, execute the following command to start the server:

```bash
go run cmd/app/main.go
```

## License

This project is licensed under the terms of the MIT license. See the LICENSE file for details.
