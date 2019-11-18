#!/usr/bin/env bash

# Deploy ping
kubectl apply -f ../deployments/ping/deployment.yml

# Deploy pong
kubectl apply -f ../deployments/pong/deployment.yml
