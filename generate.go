package apiclient

//go:generate -command protoc sh -c "protoc -I. --go_out=import_path=apiclient:. *.proto; sed -i -e s,google/protobuf,github.com/golang/protobuf/protoc-gen-go/descriptor, semantic.pb.go"
//go:generate protoc

//go:generate -command genfunc go run ./generate/gen-functions.go
//go:generate genfunc
