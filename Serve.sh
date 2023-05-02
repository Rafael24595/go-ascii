#!/bin/bash

status=$( docker container inspect -f '{{.State.Status}}' go-ascii )

if [[ $status == "exited" ]]; then
  echo "Starting container..."
  docker start go-ascii
  echo "Container started succesfully."
elif [[ $status == "running" ]]; then
  echo "Container already started"
else
  echo "Creating docker image..."
  docker build -t go-ascii .
  echo "Starting container..."
  docker run --name go-ascii -d -p 8080:8080 go-ascii	
  echo "Container started succesfully."
fi
$SHELL