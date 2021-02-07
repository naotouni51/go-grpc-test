#!/bin/sh

protoc --proto_path ./proto --go_out=plugins=grpc:./server/pb/notification notification.proto
protoc --proto_path ./proto --go_out=plugins=grpc:./client/pb/notification notification.proto