#!/usr/bin/env bash
goyacc -o comments.go comments.y && go build -o comments comments.go main.go