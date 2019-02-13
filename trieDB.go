package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const port = "4577"

func handleConnection(con net.Conn, t *trie) {
	defer func() {
		fmt.Println("Closing handle for %s \n", con.RemoteAddr().String())
		con.Close()
	}()
	fmt.Printf("Created handle for %s\n", con.RemoteAddr().String())
	ioIn := bufio.NewReader(con)
	ioOut := bufio.NewWriter(con)
	for {
		data, err := ioIn.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		data = strings.TrimSpace(string(data))
		cmds := strings.Split(data, ":")
		switch cmds[0] {
		case "FETCH":
			//code for fetch
			res, err := t.Fetch(cmds[1])
			if err != nil {
				ioOut.WriteString("false:" + err.Error())
			} else {
				ioOut.WriteString("true:" + res)
			}
		case "VINSERT":
			//returns key and inserts values

		case "TERM":
			//terminates conn
			con.Close()
		case "DEFAULT":
			//
		}
	}

}

func main() {
	fmt.Println("Initialising DB...")
	tree := Init()
	handle, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer handle.Close()
	for {
		c, err := handle.Accept()
		if err != nil {
			fmt.Println("Coneection error: ", err.Error())
			return
		}
		go handleConnection(c, tree)
	}
}
