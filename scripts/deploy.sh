#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Deploy ping
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/deployment.yml
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/service.yml

# Deploy pong
kubectl apply -f ${SCRIPT_DIR}/../deployments/pong/deployment.yml
kubectl apply -f ${SCRIPT_DIR}/../deployments/pong/service.yml
