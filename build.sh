#!/bin/bash

rm Deadication-darwin-10.6-amd64 Deadication-linux-amd64 Deadication-windows-4.0-amd64.exe

go get -v -u github.com/karalabe/xgo

xgo --targets=windows/amd64 github.com/bcvery1/Deadication
xgo --targets=darwin/amd64 github.com/bcvery1/Deadication
go build -o Deadication-linux-amd64 main.go
