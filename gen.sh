#!/bin/sh
protoc chat.proto --go_out=plugins=grpc:.
