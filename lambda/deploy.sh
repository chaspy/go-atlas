#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o hello
zip handler.zip ./hello
aws lambda update-function-code --function-name chaspy-test-function-atlas --zip-file fileb://handler.zip
