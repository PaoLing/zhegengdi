package zhegengdi

import (
	"fmt"
	// "github.com/PaoLing/zhegengdi/query"
	// "io/ioutil"
	"bytes"
	"net"
	"runtime"
)

var conMax int = 0

func handler(conn net.Conn) {

	fmt.Println(conn.RemoteAddr())
	conMax += 1
	fmt.Println(conMax)

	defer conn.Close()

	buf := make([]byte, bytes.MinRead)

	conn.Read(buf)

	fmt.Println(string(buf))

	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: text/html; charset=utf-8\r\n"))
	conn.Write([]byte("Connection: keep-alive\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte("<h1>DATABASE SERVER</h1>"))
}

func Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// connect mysql

	service := ":8080"

	laddr, err := net.ResolveTCPAddr("tcp", service)
	CheckErr(err)

	fmt.Println(laddr)

	listener, err := net.ListenTCP("tcp", laddr)
	CheckErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handler(conn)
	}
}
