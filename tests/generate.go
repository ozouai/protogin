package tests

//go:generate bash -c "cd ../protoc-gen-protogin && go build . && cd ../tests && PATH=\"$PATH:../protoc-gen-protogin\" protoc --go_out=. --protogin_out=. service.proto -I."
