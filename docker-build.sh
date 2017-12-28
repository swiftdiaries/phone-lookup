#!/bin/bash
docker build -f frontend/Dockerfile -t phone-lookup-frontend .
docker tag phone-lookup-frontend swiftdiaries/phone-lookup-frontend
docker push swiftdiaries/phone-lookup-frontend
##
docker build -f restapi/Dockerfile -t phone-lookup-api .
docker tag phone-lookup-api swiftdiaries/phone-lookup-api
docker push swiftdiaries/phone-lookup-api
##
docker build -f search/store/Dockerfile -t redis-swiftdiaries .
docker tag redis-swiftdiaries swiftdiaries/redis-swiftdiaries
docker push swiftdiaries/redis-swiftdiaries
