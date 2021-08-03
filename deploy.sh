#! /bin/bash
set -e

GOOS=linux go build -ldflags="-s -w" -tags prod -o bootstrap main.go
zip -9 code.zip bootstrap
# https://help.aliyun.com/document_detail/64204.htm?spm=a2c4g.11186623.2.10.1b92398elSkEGc#concept-2260129
fun deploy