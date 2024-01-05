#!/bin/bash

# Build Docker images
# docker buildx build -f ./Dockerfile --progress=plain --no-cache .
imagename="openexchange-otel-go"
imageversion="0.0.1"

docker buildx build -f ./devops/Dockerfile --tag andrewkandzuba/$imagename:$imageversion .