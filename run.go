package zhegengdi

import (
	"bytes"
	"fmt"
	"net"
	"zhegengdi/query"
)

func handler(conn net.Conn) {
	defer conn.Close()

	// declare an array for store type of io.Reader data
	var buf [512]byte

	// read the request data from conn into an slice
	n, err := conn.Read(buf[0:])
	CheckErr(err)

	// apply a buffer Object for write slice info
	result := bytes.NewBuffer(nil)
	result.Write(buf[0:n])

	fmt.Println(string(result.Bytes()))

	conn.Write([]byte("<h1>DATABASE SERVER</h1>"))
}

func Run() {

	// connect mysql
	query.QueryUser()

	ln, err := net.Listen("tcp", ":8080")
	CheckErr(err)
	for {
		conn, err := ln.Accept()
		CheckErr(err)

		go handler(conn)
	}
}
