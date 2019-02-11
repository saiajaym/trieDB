package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	//Returned if the key/value is not available
	ElemNotAvailable    = errors.New("Unable to find the given value for the key")
	ValueAlreadyPresent = errors.New("The value is already present")
	KeySpaceExhausted   = errors.New("Key Space Exhausted")
)

const keySpace = 62
const depth = 6
const maptable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var table = make(map[rune]int, 62)

type trie struct {
	space bool
	child [keySpace]*trie
	value *string
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

func (t *trie) InsertKey(key string) error {
	var arr = []rune(key)
	var temp = t
	for i := 0; i < depth; i++ {
		ind := table[arr[i]]
		if temp.child[ind] == nil {
			temp.child[ind] = createNode()
		}
		temp = temp.child[ind]
	}
	return nil
}

func (t *trie) Loadkeys(key []string) bool {
	for _, s := range key {
		var arr = []rune(s)
		var temp = t
		for i := 0; i < depth; i++ {
			ind := table[arr[i]]
			if temp.child[ind] == nil {
				temp.child[ind] = createNode()
			}
			temp = temp.child[ind]
		}
	}
	return true
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

	return "", nil
}

func (t *trie) Insertvalue(value string) (string, error) {
	return insert(t, 0, value)
}

func insert(t *trie, level int, value string) (string, error) {
	if level == 6 {
		if t.value != nil {
			return "", ValueAlreadyPresent
		}
		var tempStr = value
		t.value = &tempStr
		t.count++
		return "", nil
	}
	fmt.Println(level, t.value)
	for _, temp := range t.child {
		if temp.count < int64(math.Pow(62, float64(depth-level))) {
			key, err := insert(temp, level+1, value)
			if err == nil {
				t.count++
				return "", nil
			}
			if err == KeySpaceExhausted {
				continue
			}
			return "", err
		}
	}
	return "", KeySpaceExhausted
}

func main() {
	t := Init()
	fmt.Println(t.Loadkeys([]string{"asdqwe", "qwerty"}))
	fmt.Println(t.Fetch("qwerty"))
}
