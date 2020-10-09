package main

import (
	"testing"
	"testing/quick"
)

var n = 10000
func TestWithSystem(t *testing.T) {
	condition := func(a,b uint16) bool {
		return add(a,b) == add(b,a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: n})
	if err != nil {
		t.Errorf("System Error: %v\n",err)
	}
}

func TestWithItself(t *testing.T)  {
	condition := func(a,b uint16) bool {
		return add(a,b) == add(b,a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: n})
	if err != nil {
		t.Errorf("Self Error: %v\n",err)
	}
}


