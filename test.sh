#!/bin/bash
go mod download
go get -u github.com/jstemmer/go-junit-report
go test -v 2>&1 | go-junit-report > report.xml