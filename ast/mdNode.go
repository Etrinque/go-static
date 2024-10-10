package ast

import (
	"strings"
	"time"
)

type MDRoot struct {
	Meta     *MDMeta
	NodeTree []*MDNode
}

type MDNode struct {
	Type    *MDType
	Content []byte
}

type MDType struct {
	Name string
	Tag  *MDTag
}

type MDTag struct {
	Open  string
	Close string
}

// MDMeta is the relative file data, creation, author and update.
// It is NOT a part of the markdown content.
type MDMeta struct {
	Name      string
	Size      int64
	createdAt time.Time
	updatedAt time.Time
}

// MDTagMap is the representation of MD tags, mapped to the TagLabel
type MDTagMap map[string]MDTag

// Content is a section of Markdown after being parsed that retains the BlockTag Open, Close, and the Content within tags

func (md *MDTagMap) GenMap() map[string]MDTag {
	var MdTagMap = map[string]MDTag{
		"NONE":            {Open: "", Close: ""},
		"HEADING_1":       {Open: "#", Close: ""},
		"HEADING_2":       {Open: "##", Close: ""},
		"HEADING_3":       {Open: "###", Close: ""},
		"HEADING_4":       {Open: "####", Close: ""},
		"HEADING_5":       {Open: "#####", Close: ""},
		"HEADING_6":       {Open: "######", Close: ""},
		"ITALIC":          {Open: "*", Close: "*"},
		"BOLD":            {Open: "**", Close: "**"},
		"BLOCKQUOTE":      {Open: ">", Close: ""},
		"CODE":            {Open: "`", Close: "`"},
		"FENCED_CODE":     {Open: "```", Close: "```"},
		"STRIKETHROUGH":   {Open: "~~~", Close: "~~~"},
		"HORIZONTAL_RULE": {Open: "---", Close: ""},
		"HIGHLIGHT":       {Open: "==", Close: "=="},
		"TABLE":           {Open: "| ", Close: " |"},
		"DEFINITION_LIST": {Open: " : ", Close: ""},
		"UNORDERED_LIST":  {Open: "- ", Close: ""},
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

// TODO: change logic to receive only string or slice of string, convert to map method
func CheckTagMap(char string, tagmap MDTagMap) (bool, error) {
	for _, v := range tagmap {
		if strings.Contains(v.Open, char) {
			if char == v.Open {
				return true, nil
			}
			continue
		}
	}
	return false, nil
}

const (
	NONE            = ""
	HEADING_1       = "# "
	HEADING_2       = "## "
	HEADING_3       = "### "
	HEADING_4       = "#### "
	HEADING_5       = "##### "
	HEADING_6       = "###### "
	HIGHLIGHT       = "== " // 	== IMPORTANT TEXT ==
	ITALIC          = "*"   //	* text *
	BOLD            = "**"  //	** text **
	BLOCKQUOTE      = ">"
	CODE            = "`"   //	` code here `
	FENCED_CODE     = "```" // ``` Code here	```
	STRIKETHROUGH   = "~~~" //	~~~ text ~~~
	HORIZONTAL_RULE = "---"
	TABLE           = "|"  // |	Thing1	 |	Thing2	|
	LINK            = "["  //	[title](link)
	IMAGE           = "![" // 	[alt text](img link)
	FOOTNOTE        = "[^" //	[^1]
	TASK_LIST       = "-[" // -[x] -[] -[] etc.
	// UNORDERED_LIST  = "[ - + " " ] - Item - Item - Item etc.
	// ORDERED_LIST    = "[ digit + . ] 1. 2. 3. etc.
	// HEADING_ID      = "### HEADING {#custom-id}
	// DEFINITION_LIST = "term : definition
	// EMOJI           = "Who Cares
	// SUBSCRIPT       = "H~2~0	[ ~ <any character> ~ ]
	// SUPERSCRIPT     = "X^2^ 	[^ <any character> ^]
)

var tagList = []string{
	NONE,
	HEADING_1,
	HEADING_2,
	HEADING_3,
	HEADING_4,
	HEADING_5,
	HEADING_6,
	HIGHLIGHT,
	ITALIC,
	BOLD,
	BLOCKQUOTE,
	CODE,
	FENCED_CODE,
	STRIKETHROUGH,
	HORIZONTAL_RULE,
	TABLE,
	LINK,
	IMAGE,
	FOOTNOTE,
	TASK_LIST,
}
