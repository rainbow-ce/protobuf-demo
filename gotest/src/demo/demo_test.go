package demo

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func passOrDie(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestDemo(t *testing.T) {
	socket, err := net.Dial("tcp", "localhost:8000")
	passOrDie(err)
	defer socket.Close()

	// sender
	var Id int32 = 1
	var name string = "someone"
	var email string = "someone@example.com"
	sendMsg := &People{
		Name:  name,
		Id:    Id,
		Email: email,
	}

	var sendbuf []byte
	sendbuf, err = proto.Marshal(sendMsg)
	passOrDie(err)

	fmt.Println(binary.Size(sendbuf))
	fmt.Println(int64(len((sendbuf))))
	// Go is unique in that, thanks to its interfaces, it's practical to encode directly to the socket
	err = binary.Write(socket, binary.LittleEndian, sendbuf)
	passOrDie(err)

	//receiver
	receiveMsg := &People{}
	recvbuf := make([]byte, 38)
	err = binary.Read(socket, binary.LittleEndian, recvbuf)
	passOrDie(err)

	err = proto.Unmarshal(recvbuf, receiveMsg)
	passOrDie(err)

	assert.Equal(t, receiveMsg.Name, "Lysting Huang")
}
