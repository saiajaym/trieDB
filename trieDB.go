package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const port = "localhost:4577"

type writeData struct {
	value string
	ch    chan string
}

func writeHandler(t *trie, writeChan chan writeData) {
	var data writeData
	for {
		select {
		case data = <-writeChan:
			key, err := t.Insertvalue(data.value)
			if err == nil {
				data.ch <- "true:" + key
			} else {
				data.ch <- "false:" + err.Error()
			}
		}
	}
}

func connectionHandler(con net.Conn, t *trie, writeReq chan writeData) {
	defer func() {
		fmt.Printf("Closing handle for %s \n", con.RemoteAddr().String())
		con.Close()
	}()
	var w writeData
	var myWriteHandle = make(chan string, 2)
	w.ch = myWriteHandle
	fmt.Printf("Created handle for %s\n", con.RemoteAddr().String())
	ioIn := bufio.NewReader(con)
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
				con.Write([]byte("false:" + err.Error() + "\n"))
			} else {
				con.Write([]byte("true:" + res + "\n"))
			}
		case "VINSERT":
			//returns key and inserts values
			w.value = cmds[1]
			writeReq <- w
			//wait for response
			select {
			case res := <-myWriteHandle:
				con.Write([]byte(res + "\n"))
			}
		case "KINSERT":
			err := t.InsertKey(cmds[1])
			if err == nil {
				con.Write([]byte("true\n"))
			} else {
				con.Write([]byte("false\n"))
			}
		case "TERM":
			//terminates conn
			con.Close()
			break
		default:
			//default
			con.Write([]byte("invalid\n"))
		}
	}

}

func main() {
	fmt.Println("Initialising DB...")
	tree := Init()
	writeChan := make(chan writeData, 1024)
	go writeHandler(tree, writeChan)
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
		go connectionHandler(c, tree, writeChan)
	}
}
