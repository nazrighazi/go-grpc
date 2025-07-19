# go-grpc

This project is a sample implementation of a microservices architecture using gRPC for communication between services. It consists of two main services: a `products` service that manages product data and a `clients` service that provides a RESTful API to interact with the `products` service.

## Architecture

The project is divided into two main services:

*   **`products` service:** A gRPC server that connects to a PostgreSQL database to manage product information.
*   **`clients` service:** A RESTful API server built with the Echo framework that acts as a client to the `products` gRPC service.

## Features

*   gRPC communication between services.
*   RESTful API for client-side interaction.
*   PostgreSQL database for data persistence.
*   Configuration management using Viper.
*   Structured logging.

## Getting Started

### Prerequisites

*   [Go](https://golang.org/) (version 1.24.5 or later)
*   [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
*   [PostgreSQL](https://www.postgresql.org/)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/go-grpc.git
    cd go-grpc
    ```

2.  **Install Go dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Generate gRPC code:**

    ```bash
    make gen-products
    ```

4.  **Set up the database:**

    *   Make sure you have a running PostgreSQL instance.
    *   The `products` service includes database migrations. You will need to apply them to your database. The migrations are located in `products/internal/database/migrations`.

5.  **Configure the services:**

    *   **Products Service:**
        *   Copy the example configuration file: `cp products/config.example.json products/config.json`
        *   Edit `products/config.json` with your database connection details.

    *   **Clients Service:**
        *   Copy the example configuration file: `cp clients/config.example.json clients/config.json`
        *   Edit `clients/config.json` to configure the server port and other settings.

### Running the Application

1.  **Run the `products` service:**

    ```bash
    go run products/cmd/grpc/main.go
    ```

2.  **Run the `clients` service:**

    ```bash
    go run clients/cmd/api/main.go
    ```

## API Endpoints

The `clients` service exposes the following API endpoint:

*   **`GET /products`**: Retrieves a list of all products.

## gRPC Service

The `products` service exposes the following gRPC method:

*   **`GetAllProducts`**: Retrieves a list of all products.

The service is defined in the `proto/products/products.proto` file.

## Makefile Commands

*   **`make gen-products`**: Generates the Go gRPC code from the `.proto` file.
