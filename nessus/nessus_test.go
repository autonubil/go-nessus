package nessus

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	c, err := NewClientFromEnvironment(WithTrace(true))
	if err != nil {
		t.Error(err)
		return
	}
	i, err := c.Folders.List()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v", i)
}
