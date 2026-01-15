#!/bin/bash

PWD=${PWD} docker compose -f docker-compose.prod.yaml build
PWD=${PWD} docker compose -f docker-compose.prod.yaml up -d
