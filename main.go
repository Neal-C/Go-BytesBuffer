package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	// "net"
)

//practical testing

//hardly testable
// func WriteToConnection(connection net.Conn, msg []byte) error {
// 	_ , err := connection.Write(msg)
// 	return err;
// }

func WriteTo(writerCloser io.WriteCloser, message []byte) error {
	// defer writerCloser.Close();
	_ , err := writerCloser.Write(message);
	return err;
}

type BytesWriterCloser struct {
	*bytes.Buffer
	io.Closer
}

func NewBytesWriterClose() *BytesWriterCloser {
	return &BytesWriterCloser{
		Buffer: new(bytes.Buffer),
	}
}

func (self *BytesWriterCloser) Close() error {
	fmt.Println("closed")
	return nil
}

func main(){

	buffer := NewBytesWriterClose();
	buffer.Write([]byte("Hire me 🤷‍♀️"));
	buffer.WriteString("Yes, for real. It was planted");
	fmt.Println(buffer.Len());
	fmt.Println(buffer.String());
	
	if err := WriteTo(buffer, []byte("oh... 🙄")); err != nil {
		log.Fatal(err);
	}
	
	fmt.Println(buffer.String());

}