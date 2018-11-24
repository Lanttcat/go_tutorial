package baseInt

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(1, 2) == 3 {
		t.Log("Yes")
	} else {
		t.Fatal("dddd")
	}
}

