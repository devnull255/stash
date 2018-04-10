#!/bin/sh
#######################################
# Retrieve aws-go-sdk
#######################################
go get -u github.com/aws/aws-sdk-go/...
go build -o bin/application application.go
