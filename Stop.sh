#!/bin/bash

status=$( docker container inspect -f '{{.State.Status}}' go-ascii )

if [[ $status == "exited" ]]; then
  echo "Container already stopped."
elif [[ $status == "running" ]]; then
  echo "Stopping the container..."
  docker stop go-ascii
  echo "Container stopped succesfully."
else
  echo "Container does not exists."
fi
$SHELL