package jafer

import "github.com/pelletier/go-toml"

/**
create by houjingchao on 17/08/11
用于读取配置信息
[APP
name="weixin"

`jf:"#.WF.NAME"`
*/

type Toml struct {
	tree *toml.Tree
}

func NewTomlFromSource(src string) (*Toml, error) {
	tree, err := toml.Load(src)
	if err != nil {
		return nil, err
	}
	return &Toml{tree: tree}, nil
}

func NewTomlFromFilePath(path string) (*Toml, error) {
	tree, err := toml.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return &Toml{tree: tree}, nil
}

func (plugin *Toml) Prefix() string {
	return "#"
}
func (plugin *Toml) Index() int {
	return 0
}
