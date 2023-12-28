#!/bin/sh
repo=registry.aliyuncs.com/google_containers

name=registry.k8s.io/metrics-server/metrics-server:v0.6.4
src_name=metrics-server:v0.6.1
docker pull $repo/$src_name
docker tag $repo/$src_name $name
docker rmi $repo/$src_name
