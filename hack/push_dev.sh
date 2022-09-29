#!/bin/bash
docker tag app:latest docker.io/duartcs/app:dev
docker push docker.io/duartcs/app:dev

docker tag testapp:latest docker.io/duartcs/testapp:dev
docker push docker.io/duartcs/testapp:dev