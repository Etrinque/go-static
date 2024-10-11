package token

type Token string

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

var TokenList = []string{
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
