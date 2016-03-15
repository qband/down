# DownAgent

Agent to handle creation of download task on remote device.

A Web API write in Go Lang.

## feature
- add download task to local or remote downloader program
- support remote download, and move to local

## impl
- act as downloader - [1](https://github.com/golang/go/wiki/Projects#p2p-and-file-sharing)
- communicate with different downloaders

## tools
- build tool
  - gb - go get github.com/constabulary/gb/...
    - dependency management
      - gb vendor fetch --branch $branch --revision $revision $repo
      - gb vendor list
    - build
      - gb build all

- IDE
  - Intellij Community + go-lang-idea-plugin

## reference
- [GoLang Projects](https://github.com/golang/go/wiki/Projects)
  - [awesome-go#command-line](https://github.com/avelino/awesome-go#command-line)
- Command Runner
  - os/exec - [1](https://stackoverflow.com/questions/20437336/how-to-execute-system-command-in-golang-with-unknown-arguments)
  - github.com/codeskyblue/go-sh

- Remote Control
  - ssh
    - solution
      - act as ssh client
        - https://github.com/kubernetes/kubernetes/blob/b9cfab87e33ea649bdd13a1bd243c502d76e5d22/pkg/util/ssh.go
        - https://github.com/wingedpig/loom
        - https://github.com/search?l=&o=desc&q=ssh+language%3AGo&ref=advsearch&s=stars&type=Repositories&utf8=%E2%9C%93
      - communicate with a ssh client
    - lib
      - golang.org/x/crypto/ssh
  - rpc
    - gRPC
  - Web API

- orchestration
  - circuit
  - juju

- protobuf generate wrapper
  - pacman -S protobuf
  - go get -u -v github.com/golang/protobuf/proto
  - go get -u -v github.com/golang/protobuf/protoc-gen-go
  - `protoc --go_out=. *.proto`

- code examples
  - microservices
    - [go-kit/kit](https://github.com/go-kit/kit)
    - [micro/micro](https://github.com/micro/micro)
  - wego
    - cli - schachmat/wego
    - web - chubin/wttr.in
