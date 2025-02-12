
# Build the Go application
build:
	go build -o blockchain main.go

# Run the Go application
run: build
	./blockchain

# Clean the build files
clean:
	rm -f blockchain

# Build the Docker image
docker-build:
	docker build -t blockchain:latest .

# Run the Docker container
docker-run:
	docker run --rm -it blockchain:latest
