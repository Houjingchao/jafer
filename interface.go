package jafer

// create by houjingchao on 17/08/09
type (
	Init interface {
		Init()
	}
	Ready interface {
		Ready()
	}
	Destroy interface {
		Destroy()
	}
)
