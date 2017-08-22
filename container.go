package jafer

// create by houjingchao on 17/08/09
type container struct {
	//容器
	strict        bool
	orangesMap    map[string][]*Holder
	delayedFields map[string][]*DelayedField
	//plugins       map[PluginWorkTime]pluginList
}

func NewContainer() *container {
	return &container{
		orangesMap: make(map[string][]*Holder),
	}
}
