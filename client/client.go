//客户端发送封包
package main

import (
	"./../protoc"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//fmt.Println(time.Now().Unix())
	server := "127.0.0.1:5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	commandconent := &deploy.Commandconent{
		Path:    proto.String("c:/"),
		Command: proto.String("whoami"),
	}

	for i := 0; i <= 0; i++ {
		go send(i, conn, commandconent)
	}
	time.Sleep(2 * time.Second)
}

func send(i int, conn *net.TCPConn, commandconent *deploy.Commandconent) {

	fmt.Println(i)

	// 进行编码.
	data, err := proto.Marshal(commandconent)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	//fmt.Println(string(data))

	msgbuf := bytes.NewBuffer(make([]byte, 0, 1024))
	msgLen := uint32(len(data))
	binary.Write(msgbuf, binary.LittleEndian, msgLen)
	//fmt.Println(msgbuf.Bytes())
	msgbuf.Write(data)
	fmt.Println(msgbuf.Len())
	//conn.Write(msgbuf[:(msgLen + uint32(len(data)))])
	conn.Write([]byte(msgbuf.Bytes()))
}
