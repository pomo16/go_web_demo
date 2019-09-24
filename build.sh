#!/bin/bash

RUN_NAME="go_web_demo"

mkdir -p output

chmod 755 output

export GO111MODULE=on

go build -a -o output/${RUN_NAME}