package main

import (
	"flag"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to server")
)

func main() {
	flag.Parse()
}
