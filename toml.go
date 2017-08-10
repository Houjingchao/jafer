package jafer

import "github.com/pelletier/go-toml"

/**
create by houjingchao on 17/08/11
用于读取配置信息
[APP
name="weixin"
 */

type Toml struct {
	tree *toml.Tree
}

func NewTomlFromSource(src string) (*Toml, error) {
	tree, err := toml.LoadFile(src)
	if err != nil {
		return nil, err
	}
	return &Toml{tree: tree}, nil
}
