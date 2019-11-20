#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Undeploy ping
kubectl delete -f ${SCRIPT_DIR}/../deployments/ping/service.yml --ignore-not-found
kubectl delete -f ${SCRIPT_DIR}/../deployments/ping/deployment.yml --ignore-not-found

# Undeploy pong
kubectl delete -f ${SCRIPT_DIR}/../deployments/pong/service.yml --ignore-not-found
kubectl delete -f ${SCRIPT_DIR}/../deployments/pong/deployment.yml --ignore-not-found
