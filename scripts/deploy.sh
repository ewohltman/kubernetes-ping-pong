#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Deploy allow-all network policy
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/networkpolicy.yml

# Deploy ping
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/deployment.yml
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/service.yml
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/gateway.yml
kubectl apply -f ${SCRIPT_DIR}/../deployments/ping/virtualservice.yml

# Deploy pong
kubectl apply -f ${SCRIPT_DIR}/../deployments/pong/deployment.yml
kubectl apply -f ${SCRIPT_DIR}/../deployments/pong/service.yml
