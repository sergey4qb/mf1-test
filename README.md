# MF1 Test

This project is a gRPC service for user management. It is containerized using Docker and uses environment variables for configuration.


## Requirements

- [Go](https://golang.org/) (v1.23.1 or later)
- [Docker](https://www.docker.com/)

## Project Structure


## Environment Variables

Create a `.env` file in the project root with the following content:

```dotenv
# gRPC settings
GRPC_PROTOCOL=tcp
GRPC_ADDRESS=0.0.0.0
GRPC_PORT=8080
```
## Building and Running with Docker

## Generating gRPC Code from Proto Files

Before building the Docker image, you need to generate the gRPC code from your proto files. A shell script is provided for this purpose.

### Step 1: Run the Proto Generation Script

From the project root, run the following command:

```bash
./proto-generate.sh
```

This script will create the directory ./proto/pb (if it doesn't exist) and generate the gRPC code for user.proto in that directory.




### Step 2: Build the Docker Image

From the project root, run:

```bash
docker build -t mf1-test .
```
This command builds the Docker image and tags it as mf1-test.

### Step 3: Run the Docker Container

To run the container and pass environment variables from your .env file, execute:

```
docker run -p 8080:8080 --env-file .env mf1-test
```