package jafer

// create by houjingchao on 17/08/09
type container struct {
	//容器
	orangesMap map[string][]*Holder
}

func NewContainer() *container {
	return &container{
		orangesMap: make(map[string][]*Holder),
	}
}
