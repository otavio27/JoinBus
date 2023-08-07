#!/usr/bin/env bash

tag=$1

docker build -t otavio27/app_joinbus:${tag} .

docker push otavio27/app_joinbus:${tag}