package ast

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestMDMap(t *testing.T) {
	var md *MDTagMap

	tagmap := md.GenMap()
	fmt.Printf("TEST MAP: %v", tagmap)
	val := []string{
		"",
		"#",
		"##",
		"###",
		"####",
		"#####",
		"######",
		"*",
		"**",
		">",
		"1.",
		"- ",
		"`",
		"---",
		"[",
		"![",
		"| ",
		"```",
		"[^",
		" : ",
		"~~~ ",
		"-[",
		"== "}

	// key := []string{"NONE",
	// 	"H1",
	// 	"H2",
	// 	"H4",
	// 	"H6",
	// 	"BOLD",
	// 	"BLOCKQUOTE",
	// 	"CODE",
	// 	"HORIZONTAL_RULE",
	// 	"LINK",
	// 	"IMAGE",
	// 	"TABLE",
	// 	"FENCED_CODE",
	// 	"FOOTNOTE",
	// 	"STRIKETHROUGH",
	// 	"TASK_LIST",
	// 	"HIGHLIGHT",
	// }

	// TODO: rework to test for both keys:vals for:for double assert
	for _, nodeVal := range val {
		ok := assert.Contains(t, tagmap, nodeVal)
		if !ok {
			log.Errorf("map did not contian: %v", nodeVal)
		}
	}
}

func TestHTMLMap(t *testing.T) {
	// var ht *HTMLTagMap

	// typemap := ht.GenMap()

	// val := []string{}

}

func TestTreeGen(t *testing.T) {}
