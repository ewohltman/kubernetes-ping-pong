#!/usr/bin/env bash

# Undeploy ping
kubectl delete -f ../deployments/ping/deployment.yml

# Undeploy pong
kubectl delete -f ../deployments/pong/deployment.yml
