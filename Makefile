init:
	go get github.com/constabulary/gb/...
	gb vendor fetch google.golang.org/grpc
	gb vendor fetch golang.org/x/net/context
	gb vendor fetch github.com/codeskyblue/go-sh
	protoc --go_out=. src/pb/helloworld/helloworld.proto

build:
	gb build

clean:
	rm -r bin/ pkg/
