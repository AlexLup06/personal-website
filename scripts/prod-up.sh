#!/bin/bash

PWD=${PWD} docker compose -f docker-compose.prod.yml build
PWD=${PWD} docker compose -f docker-compose.prod.yml up -d
