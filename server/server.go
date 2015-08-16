package main

import (
	"./../protoc"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"

	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	netListen, err := net.Listen("tcp", ":5000")
	CheckError(err)

	defer netListen.Close()
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	// 消息缓冲
	msgbuf := bytes.NewBuffer(make([]byte, 0, 4096))
	// 数据缓冲
	databuf := make([]byte, 4096)
	// 消息长度
	length := 0
	// 消息长度uint64
	ulength := uint32(0)
	for {
		readLen, err := conn.Read(databuf)
		//fmt.Println("readLen: ", readLen)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read error")
			return
		}

		// 数据添加到消息缓冲
		readLen, err = msgbuf.Write(databuf[:readLen])
		if err != nil {
			fmt.Printf("Buffer write error: %s\n", err)
			return
		}

		// 消息分割循环
		for {
			// 消息头,需判断内容有没有
			if length == 0 && msgbuf.Len() >= 10 {
				binary.Read(msgbuf, binary.LittleEndian, &ulength)
				length = int(ulength)
				//fmt.Println("header len:", length)
			}
			if length > 0 && msgbuf.Len() >= length {
				Commandconent := &deploy.Commandconent{}
				err = proto.Unmarshal(msgbuf.Next(length), Commandconent)
				if err != nil {
					log.Fatal("unmarshaling error: ", err)
				}

				fmt.Println(Commandconent)
				execCommand(Commandconent)
				length = 0
			} else {
				//length = 0
				break
			}
		}

	}

}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func execCommand(commandconent *deploy.Commandconent) {
	fmt.Println(commandconent.GetPath(), commandconent.GetCommand())
	out, err := exec.Command(commandconent.GetCommand()).Output()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("The result is %s\n", out)
}
