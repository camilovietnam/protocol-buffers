#!/bin/zsh
protoc votingpb/voting.proto --go_out=plugins=grpc:.