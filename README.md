# Objective
The goal is create a client/server grpc communication with study purpose

# How to run ?
Just run server before then client using:   
`
go run cmd/server/server.go 
`    
`
go run cmd/client/client.go
`
# Commands/Configs 
- don't forget about env vars:
    - export GOPATH=/home/user_folder/go   
    - export GOBIN=$GOPATH/bin   
    - export PATH=$PATH:/$GOROOT:$GOPATH:$GOBIN   
    - export GOROOT=/usr/local/go   

- go mod init github.com/luciano-nascimento/grpc-go-lab
- go get google.golang.org/protobuf/cmd/protoc-gen-go
- go install google.golang.org/protobuf/cmd/protoc-gen-go
- compile protobuffer:
   - apt install -y protobuf-compiler    
   - protoc --proto_path=proto proto/*.proto --go_out=pb    
   - protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb