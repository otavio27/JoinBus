#!/usr/bin/env bash

cd ~/Projetos/JoinBus/front

quasar build -m pwa

scp -P8093 -r dist/pwa/ otavio@192.168.1.125:/home/otavio/WebSite/joinbus/
