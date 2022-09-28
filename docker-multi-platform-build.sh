#!/bin/sh

docker buildx build --platform linux/amd64,linux/arm64 --push -t coryevans2324/location-tracker:latest .
