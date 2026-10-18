// Setup:
// 	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// 	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// 	libprotoc:
// 		Linux:	 sudo apt install -y protobuf-compiler
// 		MAC:	 brew install protobuf
// 		Windows: (powershell): winget install protobuf
//	install:
//		go get google.golang.org/grpc
//		go get google.golang.org/protobuf

// Generate:
// protoc --go_out=. --go-grpc_out=. user.proto
// Somehow needs: export PATH="$PATH:$(go env GOPATH)/bin"

package main

func main1() {}
