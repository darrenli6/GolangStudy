protoc --go_out=plugins=grpc:. ./proto/*.proto


grpcurl -plaintext  localhost:9090  proto.TagService.GetTagList


