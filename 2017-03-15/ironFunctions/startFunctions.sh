#!/bin/bash

docker run -d --name functions --privileged -v ${PWD}/data:/app/data -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 iron/functions
