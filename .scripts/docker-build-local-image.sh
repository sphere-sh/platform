#!/bin/bash

# This script builds a local docker image for the platform.
#
docker build --network=host \
             -t hyperplaneNL/platform:local \
             -f .docker/Dockerfile \
             .