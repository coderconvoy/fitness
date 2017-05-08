package main

import (
	"fmt"
	"testing"
)

func Test_Load(t *testing.T) {

	l := &Location{S: "BA6 8LR"}
	err := l.Load()
	if err != nil {
		t.Fail()
	}

	fmt.Println(l.xy)
}
