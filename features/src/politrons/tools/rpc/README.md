

protoc -I login/ login/login.proto --go_out=plugins=grpc:login


where login is the directory