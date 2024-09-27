package ast

import "time"

type MDNode struct {
	BlockType *NType
	BLockTag  map[string]string
	Content   *Content
	Meta      *Meta
}

type NType struct {
	mdTag string
}

type Meta struct {
	createdAt time.Time
}

// BlockTag is the representation of MD tags, mapped to the TagLabel
type BlockTag map[string]string

// Content is a section of Markdown after being parsed that retains the BlockTag Open, Close, and the Content within tags
type Content struct {
	TagLabel     *BlockTag
	Open         string
	Close        string
	InnerContent string
}

func genMap() map[string]string {
	var MdTypeMap = map[string]string{
		"":       "NONE",            //  " "
		"#":      "H1",              //	#
		"##":     "H2",              //	##
		"###":    "H3",              //	###
		"####":   "H4",              //	####
		"#####":  "H5",              //	#####
		"######": "H6",              //	######
		"*":      "ITALIC",          //	* text *
		"**":     "BOLD",            //	** text **
		">":      "BLOCKQUOTE",      //	>
		"1.":     "ORDERED_LIST",    //	1. 2. 3. etc.
		"- ":     "UNORDERED_LIST",  //	- Item - Item - Item etc.
		"`":      "CODE",            //	` code here `
		"---":    "HORIZONTAL_RULE", //	---
		"[":      "LINK",            //	[title](link)
		"![":     "IMAGE",           // ![alt text](img link)
		"| ":     "TABLE",           // |	Thing1	 |	Thing2	|
		"```":    "FENCED_CODE",     //	```	   Code here	```
		"[^":     "FOOTNOTE",        // [^1]
		// "":"HEADING_ID",      	 // ### HEADING {#custom-id}
		" : ":  "DEFINITION_LIST", //term : definition
		"~~~ ": "STRIKETHROUGH",   // ~~~ text ~~~
		"-[":   "TASK_LIST",       //-[x] -[] -[] etc.
		// "":"EMOJI",           	 // Who Cares
		"== ": "HIGHLIGHT", // == IMPORTANT TEXT ==
		// "":"SUBSCRIPT",       	 //H~2~0
		// "":"SUPERSCRIPT"     	 //X^2^
	}
	return MdTypeMap
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
