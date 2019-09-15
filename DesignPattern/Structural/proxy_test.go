package structuraldesignpattern

import (
	"testing"
)

func TestProxyObject(t *testing.T) {
	f := &ProxyObject{}
	f.ObjDo("Run")
}
