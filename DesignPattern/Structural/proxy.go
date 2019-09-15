package structuraldesignpattern

import "log"

// IObject interface
type IObject interface {
	ObjDo(action string)
}

// Object represents real objects which proxy will delegate data
type Object struct {
	action string
}

// ObjDo implements IObject interface
func (obj *Object) ObjDo(action string) {
	log.Printf("I can, %s", action)
}

// ProxyObject proxy object
type ProxyObject struct {
	object *Object
}

// ObjDo intercept action
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "Run" {
		p.object.ObjDo(action)
	}
}
