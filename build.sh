#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o leetgo
zip leetgo-linux.zip leetgo
rm leetgo

GOOS=darwin GOARCH=amd64 go build -o leetgo
zip leetgo-macos.zip leetgo
rm leetgo

GOOS=windows GOARCH=amd64 go build -o leetgo.exe
zip leetgo-windows.zip leetgo.exe
rm leetgo.exe