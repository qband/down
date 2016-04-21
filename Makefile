init:
	go get -u github.com/constabulary/gb/...
	gb vendor fetch github.com/go-kit/kit
	gb vendor fetch github.com/buger/jsonparser
# protoc --go_out=. src/pb/helloworld/helloworld.proto

build:
	gb build
	cp src/daemon/template.json bin/config.json

run:
	bin/daemon

clean:
	rm -r bin/ pkg/ 2>/dev/null || true

dep:
	ssh vps "rm -rf down"
	rsync -azPr --exclude=.git --exclude=bin --exclude=pkg --exclude=vendor/* ../down vps:.
	ssh -t vps 'sh -c "export GOPATH=$$HOME/go; export PATH=$$HOME/go/bin:$$PATH; cd down; pwd; make init"'
	rm -rf vendor/
	rsync -azPr vps:~/down/vendor ./
