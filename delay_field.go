package jafer

import "reflect"

type DelayedField struct {
	value  reflect.Value
	field  reflect.StructField
	tag    *tag
	holder *Holder
}
