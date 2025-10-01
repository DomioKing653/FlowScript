#!/bin/bash

NAME="Flow"

if [[ "$OS" == "Windows_NT" ]]; then
    NAME="Flow.exe"
fi

go build -o "$NAME" ./src
