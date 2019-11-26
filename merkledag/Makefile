
regenerate:
	(protoc -I=. -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf -I=$(GOPATH)/src --gogo_out=. node.proto)

#regenerate:
	(protoc -I=. -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf -I=$(GOPATH)/src --gogo_out=. node.proto)

#all: node.pb.go

#node.pb.go: node.proto
    protoc --gogo_out=. --proto_path=../../../../:/usr/local/Cellar/protobuf/3.10.0/include:. $<

