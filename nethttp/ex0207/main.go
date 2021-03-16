package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	i := 0
	var m, u string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			m = strings.Fields(ln)[0]
			u = strings.Fields(ln)[1]
			fmt.Println("**method is : ", m)
			fmt.Println("**url is :", u)
		}

		if ln == "" {
			break
		}

		i++
	}

	body := fmt.Sprintf(`<html><body><h1 style="color:aliceblue">HELLO RESPONSE</h1><p><strong>Method: %s</strong></p><p><strong>Url: %s</strong></p></body></html>`, m, u)
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
