init:
	go get -u github.com/constabulary/gb/...
	gb vendor fetch github.com/go-kit/kit
	gb vendor fetch github.com/spf13/viper
# protoc --go_out=. src/pb/helloworld/helloworld.proto

build:
	gb build

clean:
	rm -r bin/ pkg/

dep:
	ssh vps "rm -rf down"
	rsync -azPr ../down vps:.
	ssh -t vps 'sh -c "export GOPATH=$$HOME/go; export PATH=$$HOME/go/bin:$$PATH; cd down; pwd; make init"'
	rsync -azPr vps:~/down/vendor ./
