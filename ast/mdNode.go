package ast

import "time"

type MDNode struct {
	Type *NType
	Meta *Meta
}

type NType struct {
	mdTag string
}

type Meta struct {
	createdAt time.Time
}

const (
	NONE            = iota
	H1              //	#
	H2              //	##
	H3              //	###
	H4              //	####
	H5              //	#####
	H6              //	######
	ITALIC          //	* text *
	BOLD            //	** text **
	BLOCKQUOTE      //	>
	ORDERED_LIST    //	1. 2. 3. etc.
	UNORDERED_LIST  //	- Item - Item - Item etc.
	CODE            //	` code here `
	HORIZONTAL_RULE //	---
	LINK            //	[title](link)
	IMAGE           // ![alt text](img link)
	TABLE           // |	Thing1	 |	Thing2	|
	FENCED_CODE     //	```	   Code here	```
	FOOTNOTE        // [^1]
	HEADING_ID      // ### HEADING {#custom-id}
	DEFINITION_LIST //term : definition
	STRIKETHROUGH   // ~~~ text ~~~
	TASK_LIST       //-[x] -[] -[] etc.
	EMOJI           // Who Cares
	HIGHLIGHT       // == IMPORTANT TEXT ==
	SUBSCRIPT       //H~2~0
	SUPERSCRIPT     //X^2^
)
