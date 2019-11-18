#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Undeploy ping
kubectl delete -f ${SCRIPT_DIR}/../deployments/ping/deployment.yml

# Undeploy pong
kubectl delete -f ${SCRIPT_DIR}/../deployments/pong/deployment.yml
