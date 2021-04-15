requirements
1. golang
2. clang-format (vscode extension)
3. vscode-proto3 (vscode extension)
4. protoc setup (brew install protobuf)
5. golang grpc (go get -u google.golang.org/grpc)
6. protoc-gen-go (go get -u github.com/golang/protobuf protoc-gen-go)
7. protoc generation file (protoc ".proto directory" --go_out=plugins=grpc:.)

usefull commands:
1. go mod init
2. go mod edit -replace
3. go mod tidy
4. go run


Creation Thoughts:
in server side
1. we can group service based on something (call it service-group)
2. in each service, create appropriate function that is needed in proto file
3. run protobuf generation file
4. implement all function that you create in proto file
5. run listener and server
in client side:
1. create object for each service-group
2. each of that object can call any function that is directly under that particular service group


new feature from grpc
1. can handle streaming data from server
2. can handle streaming data from client
3. bi-directional streaming


benefit of grpc
1. lighter, and more efficient than json because of its binary
2. supported by 11 language (go, python, java, c++, js, c#, ruby)
3. pure implementation in each language