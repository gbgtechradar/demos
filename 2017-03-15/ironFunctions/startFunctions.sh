#!/bin/bash

docker run -d --name functions --privileged -v ${PWD}/data:/app/data -p 8080:8080 iron/functions
