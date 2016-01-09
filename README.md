# DownAgent

Agent to handle creation of download task on remote device.

A Web API write in Go Lang.

## tools
- build tool
  - gb - go get github.com/constabulary/gb/...
    - dependency management
      - gb vendor fetch github.com/gin-gonic/gin
      - gb vendor list
    - build
      - gb build all

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
      - communicate with a ssh client
    - lib
      - golang.org/x/crypto/ssh
  - rpc
    - gRPC
  - Web API

- code examples
  - 
