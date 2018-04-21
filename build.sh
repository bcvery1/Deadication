#!/bin/bash

# Get the libararies
go get -v github.com/faiface/pixel
go get -v github.com/faiface/beep
go get -v github.com/faiface/mainthread
go get -v github.com/faiface/glhf
go get -v golang.org/x/image

go build main.go -o DEADication
