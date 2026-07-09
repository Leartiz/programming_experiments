package main

import (
	"fmt"
)

type Store struct {
}

func (s *Store) Set(key, val string) error {
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	return "", false
}

func (s *Store) Delete(key string) {
}

func main() {
	s := &Store{}
	_ = s.Set("k", "v")
	v, ok := s.Get("k")
	fmt.Println(v, ok)
}
