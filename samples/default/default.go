package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// START OMIT
type Name string

func (name *Name) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*name = Name(strings.TrimSpace(s))
	return nil
}

type Person struct {
	Name Name
}

func main() {
	var person Person
	_ = json.Unmarshal([]byte(`{"name": " Bill Murray "}`), &person) // ignored error
	fmt.Printf("%+v", person)
}

// END OMIT
