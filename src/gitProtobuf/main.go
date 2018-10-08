package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
)

func main() {
	user := &UserInfo{
		Message: "tiptok",
		Length:  10,
		Cnt:     20,
	}
	data, err := proto.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	tmpUser := &UserInfo{}
	_ = proto.Unmarshal(data, tmpUser)
	fmt.Println(tmpUser)
}
