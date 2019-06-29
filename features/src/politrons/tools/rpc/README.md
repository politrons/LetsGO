

protoc -I api/ api/api.proto --go_out=plugins=grpc:api


where api is the directory