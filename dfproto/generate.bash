#!/bin/bash

# exit on any error
set -e

# remove old package contents
rm -rf ./*.proto ./*.pb.go plugins core/*.proto core/*.pb.go

# download latest proto files
curl -sSL https://github.com/BenLubar/dfhack/archive/proto-docs.tar.gz | tar xz --wildcards '*.proto' --transform 's#dfhack-proto-docs/##;s#proto/##;s#library/##'

# remove proto files that aren't part of the API
rm -rf plugins/*/

# fix a warning by explicitly stating these files are protobuf version 2
sed -e 's/package /syntax = "proto2";\npackage /' -i ./*.proto plugins/*.proto

# split core proto files into a separate directory
mv Core*.proto core

# make sure there aren't any unhandled core methods
diff - <(grep '^// RPC' core/*.proto | sort) <<EOF
// RPC BindMethod : CoreBindRequest -> CoreBindReply
// RPC CoreResume : EmptyMessage -> IntMessage
// RPC CoreSuspend : EmptyMessage -> IntMessage
// RPC RunCommand : CoreRunCommandRequest -> EmptyMessage
// RPC RunLua : CoreRunLuaRequest -> StringListMessage
EOF

# split plugin proto files into subdirectories
for p in plugins/*.proto; do
	plugin=$(grep '// Plugin: ' "$p")
	plugin=${plugin##*: }
	plugin=${plugin,,}
	mkdir -p plugins/"$plugin"
	mv "$p" plugins/"$plugin"/
done

# make sure we can generate protocol buffers in Go
go get -u github.com/golang/protobuf/protoc-gen-go

# generate!
protoc --go_out=import_path=dfproto,paths=source_relative:. -I. -Icore ./*.proto
cd core
protoc --go_out=import_path=dfproto_core,paths=source_relative:. -I. -I.. ./*.proto
cd ..
for p in plugins/*/; do
	(cd "$p" && protoc --go_out="import_path=$(basename "$p"),paths=source_relative:." -I../.. -I../../core -I. ./*.proto)
done
