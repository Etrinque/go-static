package ast

type HTMLRoot struct {
	Template *Template
	Headers
}

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

type Template templ.

type HTMLBlockTag map[string]string

func (HTMLNode) GenMap()
