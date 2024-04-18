# DistributedKVStore

## Overview

DistributedKVStore is a highly available and scalable key-value storage system designed to handle data distribution across multiple nodes efficiently. It employs a static sharding mechanism using the fnv1a hashing algorithm to distribute data and integrates replication to ensure data durability and availability.

## Features

* **Static Sharding**: Uses static configurations to distribute data across several shards.
* **Replication**: Supports data replication where each shard has a primary and one or more replicas.
* **CRUD Operations**: Enables basic CRUD (Create, Read, Update, Delete) operations on stored key-value pairs.
* **Fault Tolerance**: Includes mechanisms to handle node failures, maintaining continuous system operations.

## Getting Started

### Prerequisites

* Golang 1.15 or higher
* curl (for interacting with the HTTP server)
* Linux environment (recommended for executing bash scripts)

### Installation
Clone the repository to your local machine:

```bash
git clone [repository-url]
cd DistributedKVStore
```

### Configuration
Configure your shards using the sharding.toml file. Below is a sample configuration:

```bash
[[shard]]
name = "Moscow"
idx = 0
address = "127.0.0.1:8080"
replicas = ["127.0.0.1:8081"]

[[shard]]
name = "Minsk"
idx = 1
address = "127.0.0.1:8082"
replicas = ["127.0.0.1:8083"]

[[shard]]
name = "Kiev"
idx = 2
address = "127.0.0.1:8084"
replicas = ["127.0.0.1:8085"]

[[shard]]
name = "Tashkent"
idx = 3
address = "127.0.0.1:8086"
replicas = ["127.0.0.1:8087"]
```

### Running the System
To start the system with all configured shards and their replicas, execute:


```bash
./launch.sh
```
This script will launch multiple service instances based on the sharding configuration.

### Populating the Store with Data
To populate the store with random data, run:

```bash
./populate.sh
```

## API Usage
Interact with the DistributedKVStore using the following HTTP API endpoints:

* PUT /set - Stores a key-value pair.
```bash
curl -X PUT "http://127.0.0.1:8080/set?key=exampleKey&value=exampleValue"
```

* GET /get - Retrieves the value by key.
```bash
curl "http://127.0.0.1:8080/get?key=exampleKey"
```

* DELETE /delete - Removes a key-value pair.
```bash
curl -X DELETE "http://127.0.0.1:8080/delete?key=exampleKey"
```

* POST /update - Updates the value of an existing key.
```bash
curl -X POST "http://127.0.0.1:8080/update?key=exampleKey&value=newValue"
```