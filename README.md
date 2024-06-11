# DistributedKVStore

## Overview

DistributedKVStore is a highly available and scalable key-value storage system designed to handle data distribution across multiple nodes efficiently. It employs a static sharding mechanism using the fnv1a hashing algorithm to distribute data and integrates replication to ensure data durability and availability.

## Features

* **Static Sharding**: Uses static configurations to distribute data across several shards.
* **Replication**: Supports data replication where each shard has a primary and one or more replicas.
* **CRUD Operations**: Enables basic CRUD (Create, Read, Update, Delete) operations on stored key-value pairs.

## Getting Started

### Prerequisites

* Golang 1.15 or higher
* curl (for interacting with the HTTP server)
* Linux environment (recommended for executing bash scripts)

### Installation
Clone the repository to your local machine:

```bash
git clone git@github.com:RE110010100/Distributed_KV_Store.git
cd DistributedKVStore
```

### Configuration
Configure your shards using the sharding.toml file. Below is a sample configuration:

```bash
[[shard]]
name = "Shard 1"
idx = 0
address = "192.168.0.1:8080"
replicas = ["192.168.0.1:8081"]

[[shard]]
name = "Shard 2"
idx = 1
address = "192.168.0.1:8082"
replicas = ["192.168.0.1:8083"]

#add more shards here
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

## Benchmarking 

### Overview 

The benchmarking tool provided with DistributedKVStore is designed to evaluate the performance of the system under various load conditions. It measures the throughput (queries per second, QPS) for write and read operations across the configured shards.

### Prerequisites

Ensure you have Go installed on your machine as the benchmark tool is a Go program. This tool also assumes that the DistributedKVStore is already running and accessible.

### Setup and Execution

1. **Download the benchmark tool:** Clone the repository if you haven't already, and navigate to the cmd/bench directory.

```bash
git clone git@github.com:RE110010100/Distributed_KV_Store.git
cd DistributedKVStore/cmd/bench/
```

2. **Build the benchmark tool:** Compile the program to ensure it's ready to run.

```bash
go build main.go
```

3. **Make sure that an instance of the DistributedKVStore is running:** To ensure that the DistributedKVStore app is running before it's componenets can be benchmarked. 

4. **Run the benchmark:** Execute the benchmark program. You can specify the address of the server instance, the number of iterations for write and read operations, and the concurrency level (number of goroutines running in parallel).

```bash
./main -addr="localhost:8080" -iterations=1000 -read-iterations=100000 -concurrency=10
```

* **addr**: The HTTP host and port for the instance that is being benchmarked.
* **iterations**: The number of iterations for writing.
* **read-iterations**: The number of iterations for reading.
* **concurrency:** How many goroutines to run in parallel during the tests.

Adjust the parameters based on your testing requirements and the configuration of your system.

### Interpreting Results

The benchmark tool will output the average execution time, the QPS, and the minimum and maximum latencies for both write and read operations. This data can help identify performance bottlenecks and the scalability of the system under load.

Use this information to tune your DistributedKVStore setup or to understand the limits of the current configuration under simulated load conditions.



