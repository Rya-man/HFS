package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	// "sort"
	// "os"
	// "strings"
)
func getLinesChannel(f io.ReadCloser) <- chan string{
	out:= make(chan string,1)
	go func ()  {
		defer f.Close()
		defer close(out)
		str := ""
		for {
			var buf = make([]byte,8)
			n,err := f.Read(buf)
			if err != nil{
				break
			}
			buf = buf[:n]
			if j:= bytes.IndexByte(buf,'\n'); j!=-1{
				str += string(buf[:j])
				buf = buf[j+1:]
				out <- str
				str = ""
			}
			str += string(buf)
			
		}
		if len(str)!=0{
				out <- str
			}
	}()
	return out
}

func main(){

	listener, err := net.Listen("tcp",":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	for{
		conn,err := listener.Accept()
		if err!= nil {
			log.Fatal("error", "error", err)
		}

		for line:= range getLinesChannel(conn){
			fmt.Printf("Read: %s\n",line)
		}
	}
	
}
