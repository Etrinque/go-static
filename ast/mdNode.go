package ast

import (
	"time"
)

type MDRoot struct {
	Meta     *MDMeta
	NodeTree []*MDNode
}

type MDNode struct {
	Tags    *MDTag
	Content []byte
}

type MDType struct {
	Name string
	Tag  *MDTag
}

type MDTag struct {
	open  string
	close string
}

// MDMeta is the relative file data, creation, author and update.
// It is NOT a part of the markdown content.
type MDMeta struct {
	Name      string
	Size      int64
	createdAt time.Time
	updatedAt time.Time
}

// BlockTag is the representation of MD tags, mapped to the TagLabel
type MDTagMap map[string]MDTag

// Content is a section of Markdown after being parsed that retains the BlockTag Open, Close, and the Content within tags

func (MDTagMap) GenMap() map[string]MDTag {
	var MdTagMap = map[string]MDTag{
		"NONE":            {open: "", close: ""},
		"HEADING_1":       {open: "#", close: ""},
		"HEADING_2":       {open: "##", close: ""},
		"HEADING_3":       {open: "###", close: ""},
		"HEADING_4":       {open: "####", close: ""},
		"HEADING_5":       {open: "#####", close: ""},
		"HEADING_6":       {open: "######", close: ""},
		"ITALIC":          {open: "*", close: "*"},
		"BOLD":            {open: "**", close: "**"},
		"BLOCKQUOTE":      {open: ">", close: ""},
		"CODE":            {open: "`", close: "`"},
		"FENCED_CODE":     {open: "```", close: "```"},
		"STRIKETHROUGH":   {open: "~~~", close: ""},
		"HORIZONTAL_RULE": {open: "---", close: ""},
		"HIGHLIGHT":       {open: "==", close: "=="},
		"TABLE":           {open: "| ", close: " |"},
		"DEFINITION_LIST": {open: " : ", close: ""},
		"UNORDERED_LIST":  {open: "- ", close: ""},
		// "ORDERED_LIST":    {open: fmt.Sprintf("%d."), close: ""},
		// TODO: add link/IMG support
		// "LINK":  {open: fmt.Sprintf("[%s](%s)"), close: ""},
		// "IMAGE": {open: fmt.Sprintf("![%s](%s)"), close: ""},
		//
		// "FOOTNOTE":        {open: fmt.Sprintf("[^%d]"), close: ""},
		// "TASK_LIST":       {open: fmt.Sprintf("-[%s]"), close: ""},
	}
	return MdTagMap
}

const (
	NONE            = ""
	H1              = "#"
	H2              = "##"
	H3              = "###"
	H4              = "####"
	H5              = "#####"
	H6              = "######"
	ITALIC          = "*"  //	* text *
	BOLD            = "**" //	** text **
	BLOCKQUOTE      = ">"
	CODE            = "`" //	` code here `
	HORIZONTAL_RULE = "---"
	LINK            = "["   //	[title](link)
	IMAGE           = "!["  // 	[alt text](img link)
	TABLE           = "|"   // |	Thing1	 |	Thing2	|
	FENCED_CODE     = "```" // ``` Code here	```
	FOOTNOTE        = "[^"  //	[^1]
	STRIKETHROUGH   = "~~~" //	~~~ text ~~~
	TASK_LIST       = "-["  // -[x] -[] -[] etc.
	HIGHLIGHT       = "=="  // 	== IMPORTANT TEXT ==
	// UNORDERED_LIST  = "[ - + " " ] - Item - Item - Item etc.
	// ORDERED_LIST    = "[ digit + . ] 1. 2. 3. etc.
	// HEADING_ID      = "### HEADING {#custom-id}
	// DEFINITION_LIST = "term : definition
	// EMOJI           = "Who Cares
	// SUBSCRIPT       = "H~2~0	[ ~ <any character> ~ ]
	// SUPERSCRIPT     = "X^2^ 	[^ <any character> ^]
)
