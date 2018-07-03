//go:generate sh -ec "rm -rf *.proto *.pb.go"
//go:generate sh -ec "curl -sSL https://github.com/DFHack/dfhack/archive/develop.tar.gz | tar xz --wildcards '*.proto' --transform 's#.*/##' && chmod -x *.proto && rm -f isoworldremote.proto && sed -e $'s/package /syntax = \x22proto2\x22;\\npackage /' -i *.proto"
//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate sh -ec "protoc --go_out=import_path=dfproto,paths=source_relative:. *.proto"

package dfproto
