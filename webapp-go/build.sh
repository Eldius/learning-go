#!/bin/bash

echo "- Building from module script..."
yarn install && \
    go clean && go test && go build
