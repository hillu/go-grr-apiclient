package apiclient

//go:generate -command protoc sh -c "protoc -I. --go_out=import_path=apiclient:. *.proto; sed -i -e s,google/protobuf,github.com/golang/protobuf/protoc-gen-go/descriptor, semantic.pb.go"
//go:generate protoc

//go:generate -command genfunc ./generate/gen-functions.sh

//go:generate genfunc SearchClients get /api/clients

//go:generate genfunc AddClientsLabels    post-simple /api/clients/labels/add
//go:generate genfunc RemoveClientsLabels post-simple /api/clients/labels/remove
//go:generate genfunc ListClientsLabels   get-simple  /api/clients/labels

//go:generate genfunc ListArtifacts   get         /api/artifacts
//go:generate genfunc DeleteArtifacts post-simple /api/artifacts/delete
//go:generate genfunc UploadArtifact  post-simple /api/artifacts/upload
