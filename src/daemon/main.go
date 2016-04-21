package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read configuration file
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)
	config, err := ioutil.ReadFile(path.Join(dir, "config.json"))
	check(err)
	config, _, _, _ = jsonparser.Get(config, "downloader")
	fmt.Println(string(config))
}
