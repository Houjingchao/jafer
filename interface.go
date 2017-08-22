package jafer

// create by houjingchao on 17/08/09

type Obj interface{}

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

type Sharer interface {
	Share() interface{}
}
