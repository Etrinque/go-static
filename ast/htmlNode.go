package ast

import (
	"fmt"

	"github.com/a-h/templ"
)

type HTMLRoot struct {
	Template *Template
	Headers
}

type Headers struct {
	Tag  string
	Meta *MDMeta
}

type HTMLNode struct {
	TagType  *HTMLType
	BlockTag *HTMLTagMap
	Content  []byte
}

type HTMLType struct {
	Name string
	Tags *HTMLTag
}

type HTMLTag struct {
	open  string
	close string
}

type Template *templ.Component

type HTMLTagMap map[string]HTMLTag

var link string
var imgLink string
var imgTxt string

// TODO: Add support for HMTL Tables. Maybe Struct of Structs of Elements?

func (h *HTMLTagMap) GenMap() *HTMLTagMap {
	HTMLTagMap := HTMLTagMap{
		"NONE":            {open: "", close: ""},
		"HEADER":          {open: "<header>", close: "</header>"},
		"META":            {open: "<meta>", close: "</meta>"},
		"HEADING_1":       {open: "<h1>", close: "</h1>"},
		"HEADING_2":       {open: "<h2>", close: "</h2>"},
		"HEADING_3":       {open: "<h3>", close: "</h3>"},
		"HEADING_4":       {open: "<h4>", close: "</h4>"},
		"HEADING_5":       {open: "<h5>", close: "</h5>"},
		"HEADING_6":       {open: "<h6>", close: "</h6>"},
		"PARAGRAPH":       {open: "<p>", close: "</p>"},
		"ORDERED_LIST":    {open: "<ol>", close: "</ol>"},
		"LIST_ITEM":       {open: "<li>", close: "</li>"},
		"UNORDERED_LIST":  {open: "<ul>", close: "</ul>"},
		"BLOCKQUOTE":      {open: "<blockquote>", close: "</blockquote>"},
		"CODE_BLOCK":      {open: "<pre><code>", close: "</code></pre>"},
		"HORIZONTAL_RULE": {open: "<hr/>", close: ""},
		"BOLD":            {open: "<b>", close: "</b>"},
		"ITALIC":          {open: "<i>", close: "</i>"},
		"STRIKETHROUGH":   {open: "<del>", close: "</del>"},
		"LINK":            {open: fmt.Sprintf("<a href=\" %v \">", link), close: "</a>"},
		"IMAGE":           {open: fmt.Sprintf("<img src=\" %v \" alt=\" %v \">", imgLink, imgTxt), close: ""},
	}
	return &HTMLTagMap
}
