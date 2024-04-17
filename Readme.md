# Zocket Key-Value Store

Zocket Key-Value Store is a Go Lang application that implements a distributed key-value store using [HashiCorp Consul](https://developer.hashicorp.com/consul/docs) for storing data and service discovery.

## Overview

The application allows users to store and retrieve key-value pairs via HTTP API endpoints. It utilizes HashiCorp Consul's Key-Value Store for data storage and replication across the cluster. Service discovery is also handled by Consul, enabling dynamic addition and removal of nodes from the cluster.

## Features

- **Distributed Key-Value Store:** Nodes in the cluster can store and retrieve key-value pairs.
- **Membership Management:** HashiCorp Memberlist is used for managing cluster membership and communication between nodes.
- **Service Discovery:** Consul is used for service discovery, allowing new nodes to join the cluster dynamically.

## Installation

To build and run the Zocket Key-Value Store application, follow these steps:

1. **Ensure you have Go installed on your system.**
2. **Ensure you have installed Consul [CLick here to go to download page](https://developer.hashicorp.com/consul/install)**
3. **Start the Consul by running the command.** 
     ```bash
     consul agent -dev
     ```
4. **Clone the repository:**

   ```bash
   git clone https://github.com/KhalithEngineer/zocket-key-value-store
   ```
5. **Change into the project directory:**

    ```bash
    cd zocket-key-value-store
    ```

6. **Build the application:**

    ```bash
    go build -o zocket-key-value-store ./cmd/server
    ```
8. **Export the envs:**

    ```bash
    export $(cat .env | xargs)
    ```

7. **Run the application:**

    ```bash
    ./zocket-key-value-store
    ```
8. **GET/PUT/DELETE can be done by REST-API:**

    ```bash
    curl --location 'http://localhost:8086/put' \
        --header 'Content-Type: application/json' \
        --data '{
        "key": "iAmKey",
        "value": "ItsmeTheValue"}'
    ```
     ```bash
    curl --location 'http://localhost:8080/get' \
        --header 'Content-Type: application/json' \
        --header 'Accept: application/json' \
        --data '{"key": "iAmKey"}'
    ```
     ```bash
        curl --location 'http://localhost:8080/delete' \
        --header 'Content-Type: application/json' \
        --data '{
        "key": "iAmKey"}'
    ```