#!/bin/bash
docker build -f frontend/Dockerfile -t phone-lookup-frontend .
docker build -f restapi/Dockerfile -t phone-lookup-api .