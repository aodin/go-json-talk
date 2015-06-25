package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// START OMIT
type PrimaryKey struct{ ID int }

func (pk PrimaryKey) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(pk.ID)), nil
}

type Person struct {
	PrimaryKey
	Name string
}

func main() {
	b, _ := json.Marshal(PrimaryKey{ID: 1}) // ignored error
	fmt.Printf("%s\n", b)
	b, _ = json.Marshal(Person{PrimaryKey{ID: 1}, "HAL"}) // ignored error
	fmt.Printf("%s\n", b)
}

// END OMIT
