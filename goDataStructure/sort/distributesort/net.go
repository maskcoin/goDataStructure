package distributesort

import (
	"bufio"
	"net"
)


func NetworkWrite(addr string, in <-chan int) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer ln.Close()
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSync(writer, in)
	}()
}


func NetworkRead(addr string) <-chan int {
	out := make(chan int)

	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		r := ReadSource(bufio.NewReader(conn), -1)
		for v := range r {
			out<-v
		}
		close(out)
	}()
	return out
}
