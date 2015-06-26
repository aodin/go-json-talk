package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BEGIN OMIT
type Key int64

// Allow the key to only be set once
func (key *Key) UnmarshalJSON(data []byte) error {
	if *key != 0 {
		return fmt.Errorf("A key cannot be changed after it is set")
	}
	num, err := strconv.ParseInt(string(data), 10, 64)
	*key = Key(num)
	return err
}

func main() {
	var key Key
	_ = json.Unmarshal([]byte(`23`), &key) // ignored error
	fmt.Println(key)

	// Attempt to overwrite
	fmt.Println(json.Unmarshal([]byte(`24`), &key))
	fmt.Println(key)
}
