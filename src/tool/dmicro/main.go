package main

import "github.com/tiptok/OFAppTest/src/tool/dmicro/cmd"

func main(){
	cmd.Init(
		cmd.Name("dmicro"),
		cmd.Version("1.0.0"),
		cmd.Description("A microservice toolkit"),
		)
}
