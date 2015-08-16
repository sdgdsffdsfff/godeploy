package main

import (
	"log"
	// 辅助库
	"./protoc"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	command := &deploy.Commandconent{
		Path:    proto.String("c:/"),
		Command: proto.String("whoami"),
	}
	// 进行编码
	data, err := proto.Marshal(command)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(string(data))
	// 进行解码
	decode := &deploy.Commandconent{}
	err = proto.Unmarshal(data, decode)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(command, "\n", data, "\n", decode, "\n", decode.GetPath())
}
