#!/bin/bash

echo "- Building from module script..."
go clean && go test && go build
