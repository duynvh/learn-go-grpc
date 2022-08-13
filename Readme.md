protoc -I=proto --go_out=pb proto/*.proto

protoc -I=proto proto/*.proto --go_out=plugins=grpc:pb