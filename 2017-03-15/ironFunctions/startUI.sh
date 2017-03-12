#!/bin/bash
docker run -d --link functions:api -p 4000:4000 -e "API_URL=http://api:8080" iron/functions-ui
