#!/bin/bash
docker build .. -f ../Dockerfile --target app -t app:latest --platform linux/amd64
docker build .. -f ../Dockerfile --target test-app -t testapp:latest --platform linux/amd64