# Zocket Key-Value Store

Zocket Key-Value Store is a Go Lang application that implements a distributed key-value store using HashiCorp Memberlist for membership management and communication between nodes.

## Overview

The application is designed to run in a distributed environment where multiple nodes can join a cluster and store key-value pairs. Each node in the cluster communicates with other nodes using HashiCorp Memberlist, ensuring fault tolerance and scalability. Consul is used for service discovery, enabling new nodes to join the cluster dynamically.

## Features

- **Distributed Key-Value Store:** Nodes in the cluster can store and retrieve key-value pairs.
- **Membership Management:** HashiCorp Memberlist is used for managing cluster membership and communication between nodes.
- **Service Discovery:** Consul is used for service discovery, allowing new nodes to join the cluster dynamically.

## Installation

To build and run the Zocket Key-Value Store application, follow these steps:

1. **Ensure you have Go installed on your system.**
2. **Clone the repository:**

   ```bash
   git clone <repository-url>
3. **Change into the project directory:**

    ```bash
    cd zocket-key-value-store
    ```

4. **Build the application:**

    ```bash
    go build -o zocket-key-value-store ./cmd/server
    ```

5. **Run the application:**

    ```bash
    ./zocket-key-value-store
    ```