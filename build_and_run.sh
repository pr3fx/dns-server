#!/bin/sh
set -e # Exit early if any commands fail
BUILD_DIR=$(dirname "$0")/build
EXECUTABLE_DIR=$BUILD_DIR/bin

(
    mkdir -p $EXECUTABLE_DIR # Ensure build and executable directories exists
    cd $(dirname "$0")       # Build commands run from this script's directory
    go build -o $EXECUTABLE_DIR/dns-server-go cmd/main.go
)

exec $EXECUTABLE_DIR/dns-server-go "$@"
