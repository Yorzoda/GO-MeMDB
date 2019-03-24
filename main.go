package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer ls.Close()
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)

	}
}
func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\n In memory DataBase\n"+
		"Use:\n"+
		"Set key value\n"+
		"Get key\n"+
		"DEl key\n"+
		"Example:\n"+
		"Set ooo offers\n"+
		"Get ooo offers\n\n\n")
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch fs[0] {
		case "Get":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "Set":
			if len(fs) != 3 {
				fmt.Println(conn, "Expected value")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "Del":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Println(conn, "Invalid comand")
		}
	}
}
