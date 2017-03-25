#!/bin/sh

path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go|sh)
        gofmt -w $path
	go build .
	go install ./show_slide
        ;;
esac

