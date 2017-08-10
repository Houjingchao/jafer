package jafer

import "strings"

// create by houjingchao on 17/08/09

type tag struct {
	auto   bool
	depend bool
	name   string
	path   string
	prefix string
}

//根据tag 生成 tag struct
func parseTag(tag string) *tag {
	t := &tag{}

	if tag == "auto" {
		t.depend = true
		t.auto = true
		return t
	}

	if len(tag) < 1 {
		panic("tag length <1 ")
		return t
	}

	if strings.Contains(tag, ".") {
		t.depend = false
		t.prefix = tag[:strings.Index(tag, ".")]
		t.path = tag[strings.Index(tag, ".")+1:]
		return t
	}

	t.depend = true
	t.name = tag
	return t
}
