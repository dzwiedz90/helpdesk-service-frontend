#!/bin/bash

go build -o helpdesk-service-frontend main.go && ./helpdesk-service-frontend | tee logs/console.log