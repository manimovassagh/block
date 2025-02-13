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
