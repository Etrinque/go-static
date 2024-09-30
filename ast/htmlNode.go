package ast

type HTMLNode struct {
	TagType  *HTMLType
	BlockTag *HTMLBlockTag
	Content  *Content
}

type HTMLType struct {
	Name string
	Tag  string
}

type Headers struct {
	Tag  string
	Meta *MDMeta
}

type HTMLBlockTag map[string]string

func (HTMLNode) GenMap()
