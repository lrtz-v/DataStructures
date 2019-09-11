package creationaldesignpattern

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {

	s := NewSingleton()

	s["this"] = "that"

	s2 := NewSingleton()

	fmt.Println("This is ", s2["this"])
}
