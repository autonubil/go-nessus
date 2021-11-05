package nessus

import (
	"fmt"
	"testing"
)

func TestStatus(t *testing.T) {
	c, err := NewClientFromEnvironment(WithTrace(true))
	if err != nil {
		t.Error(err)
		return
	}
	params := ScansListParams{}
	i, err := c.Scans.List(&params)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v", i)
}
