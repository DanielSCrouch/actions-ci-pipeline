#!/bin/bash
docker tag app:latest docker.io/duartcs/app:pro
docker push docker.io/duartcs/app:pro

docker tag testapp:latest docker.io/duartcs/testapp:pro
docker push docker.io/duartcs/testapp:pro