#!/bin/sh

# # Enable BuildKit for more efficient builds
# export DOCKER_BUILDKIT=1

# Define variables
IMAGE_NAME="stylize4"
IMAGE_TAG="1.0"
CONTAINER_NAME="stylize_container4"
PORT="8080"

# Step 1: Build the Docker image and give it a tag and name
docker image build -f Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Docker image built successfully and named successfully."
else
  echo "Docker image build failed."
  exit 1
fi

# Step 2: Run the Docker container
echo "Running the Docker container..."
docker run -d -p ${PORT}:${PORT} --name ${CONTAINER_NAME} ${IMAGE_NAME}:${IMAGE_TAG}

# Check if the container is running
if [ $? -eq 0 ]; then
  echo "Docker container is running."
else
  echo "Failed to run the Docker container."
  exit 1
fi

echo "Container ${CONTAINER_NAME} is running and accessible on port ${PORT}."
