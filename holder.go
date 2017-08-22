package jafer

import "reflect"

//create by houjingchao on 17/08/09
type Holder struct {
	Ignore bool
	Obj    Obj
	Type reflect.Type
	PointerType reflect.Type
	Value  reflect.Value

}
