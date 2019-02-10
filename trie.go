package main

import (
	"errors"
	"fmt"
)

var (
	//Returned if the key/value is not available
	ElemNotAvailable = errors.New("Unable to find the given value for the key")
)

const keySpace = 62
const depth = 6
const maptable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var table = make(map[rune]int, 62)

type trie struct {
	space bool
	child [keySpace]*trie
	value string
	count int64
}

func buildMap() {
	for ind, k := range maptable {
		table[k] = ind
	}
	fmt.Println(table)
}

//inits trie, call this for the first time
func Init() *trie {
	buildMap()
	return new(trie)
}

func createNode() *trie {
	var newNode = new(trie)
	newNode.space = true
	return newNode
}

func (t *trie) InsertKey(key string) string {
	var arr = []rune(key)
	var temp = t
	for i := 0; i < depth; i++ {
		ind := table[arr[i]]
		fmt.Println(temp, ind)
		if temp.child[ind] == nil {
			temp.child[ind] = createNode()
		}
		temp = temp.child[ind]
	}
	return "suc"
}

func (t *trie) Fetch(key string) (string, error) {
	var temp = t
	var arr = []rune(key)
	for i := 0; i < depth; i++ {
		ind := table[arr[i]]
		if temp.child[ind] == nil {
			return "", ElemNotAvailable
		}
		temp = temp.child[ind]
	}

	return "", temp.value
}

func main() {
	t := Init()
	fmt.Println(t.InsertKey("asdqwe"))
	fmt.Println(t.Fetch("asdqwe"))
}
