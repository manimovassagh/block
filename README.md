# Simple Blockchain in Go

This project implements a simple blockchain with proof of work and mining in Go.

## Structure

- `main.go`: Contains the main logic for creating and managing the blockchain.
- `models/block.go`: Contains the `Block` struct and its methods.
- `Makefile`: Contains commands to build, run, and clean the project, as well as Docker commands.
- `Dockerfile`: Defines the Docker image for the project.

## How to Run

1. Ensure you have Go installed on your machine.
2. Clone the repository or download the project files.
3. Navigate to the project directory.

### Using Makefile

To build and run the application using the Makefile, use the following commands:

```sh
make build   # Build the application
make run     # Run the application
make clean   # Clean the build files
```

### Using Docker

To build and run the application using Docker, use the following commands:

```sh
make docker-build   # Build the Docker image
make docker-run     # Run the Docker container
```

## Example Output




0000000000000000000
18b7e6456b26ae8cc31ba30461f3986cb720a7c92dec3
Base is this
We have Three entities 
genesis block  and normal block and next block
any block has to have difficulty 0000 s on start of hash (depends on difficulty could be 1 or 4 or 10 etc )
nounce would be the number that create this proper start 0000 numbers on begin of hash 
that means we need a hash that data + previous hash and put some randome noance at beging and do it until we become 
for example 0000 at begin of hash
in that case the block is mined
this is mining
