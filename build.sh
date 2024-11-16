#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o leetgo

GOOS=darwin GOARCH=amd64 go build -o leetgo-darwin


GOOS=windows GOARCH=amd64 go build -o leetgo.exe
