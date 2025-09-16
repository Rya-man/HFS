package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	x:="localhost:42069"
	udpaddr,err:= net.ResolveUDPAddr("udp",x)
	if err!=nil {
		return
	}
	conn,err:= net.DialUDP("udp",nil,udpaddr)
	if err!=nil {
		return
	}
	defer conn.Close()

	fmt.Printf("Sending to %s Type your message and press enter.",x)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		msg,err:= reader.ReadString('\n')
		if err!=nil{
			continue
		}
		_,err = conn.Write([]byte(msg))
		if err!=nil{
			continue
		}

		fmt.Printf("message sent: %s",msg)
	}
}