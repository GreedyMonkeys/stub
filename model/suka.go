package model

import "encoding/json"

type Suka struct {
	Name    string
	Surname string
}

func (s *Suka) ToJson() ([]byte, error) {
	return json.Marshal(s)
}

func add(a int, b int) int {
	return a + b
}
