package main

import (
	"github.com/tiptok/OFAppTest/src/tool/gencode/cmd"
)

func main(){
	cmd.Init(
		cmd.Name("gencode"),
		cmd.Version("0.0.1"),
		cmd.Description("a tool gen code"),
	)
}


