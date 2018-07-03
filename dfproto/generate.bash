#!/bin/bash

# exit on any error
set -e

# remove old package contents
rm -rf *.proto *.pb.go

# download latest proto files
curl -sSL https://github.com/DFHack/dfhack/archive/develop.tar.gz | tar xz --wildcards '*.proto' --transform 's#.*/##'

# remove execute permission (some of the DFHack proto files are weird)
chmod -x *.proto

# delete isoworldremote (it conflicts with RemoteFortressReader)
rm -f isoworldremote.proto

# fix a warning by explicitly stating these files are protobuf version 2
sed -e 's/package /syntax = "proto2";\npackage /' -i *.proto

# make sure we can generate protocol buffers in Go
go get -u github.com/golang/protobuf/protoc-gen-go

# generate!
protoc --go_out=import_path=dfproto,paths=source_relative:. *.proto
