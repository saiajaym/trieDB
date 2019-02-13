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
var powTable []int64

type trie struct {
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
	t := new(trie)
	t.count = 0
	return t
}

func createNode() *trie {
	var newNode = new(trie)
	newNode.count = 0
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

	return *temp.value, nil
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
	for loc, temp := range t.child {
		if temp != nil && temp.count < int64(math.Pow(62, float64(depth-level))) {
			key, err := insert(temp, level+1, value)
			if err == nil {
				key = string(maptable[loc]) + key
				t.count++
				return key, nil
			}
			if err == KeySpaceExhausted || err == ValueAlreadyPresent {
				continue
			}
			return "", err
		}
	}
	return "", KeySpaceExhausted
}

/*func main() {
	t := Init()
	fmt.Println(t.Loadkeys([]string{"asdqwe", "qwerty"}))
	fmt.Println(t.Insertvalue("hola bitch"))
	fmt.Println(t.Fetch("asdqwe"))
	fmt.Println(t.Insertvalue("test me out"))
	fmt.Println(t.Fetch("qwerty"))
	s := "asfdf"
	for i := 0; i < 62; i++ {
		fmt.Println("Insering", s+string(maptable[i]))
		t.InsertKey(s + string(maptable[i]))
		fmt.Println(t.Insertvalue("hre"))
	}
	//fmt.Println(t.InsertKey("asfdgg"))
	//fmt.Println(t.InsertKey("asfdgh"))
	fmt.Println(t.Insertvalue("BABYYYY"))
	fmt.Println(t.Fetch("asfdgg"))
	s = s + "a"
}*/
