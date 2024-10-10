package ast

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMDMap(t *testing.T) {
	var md *MDTagMap

	tagmap := md.GenMap()
	t.Logf("TEST MAP: %v", tagmap)
	testval := []string{
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

	// TODO: rework to test for both keys:vals for:for double assert
	tags := maps.Values[MDTagMap](tagmap)
	for v := range tags {
		tag := v.Open
		for _, test := range testval {
			ok := assert.Contains(t, tag, test)
			if !ok {
				t.Errorf("map did not contian: %v", test)
			}
			return
		}
	}
}

func TestCheckTagmap(t *testing.T) {
	var md *MDTagMap
	var tt = []string{"**", "*", "***", "a", "# ", "### "}

	tagmap := md.GenMap()
	if tagmap == nil {
		t.Errorf("error initializing map got=%v", tagmap)
	}

	for _, s := range tt {
		ok, err := CheckTagMap(s, tagmap)
		if !ok {
			t.Logf("failure to match: %s, got=%s", s, tagmap)
		}
		if err != nil {
			t.Error(err)
		}
	}
}

func TestHTMLMap(t *testing.T) {
	// var ht *HTMLTagMap

	// typemap := ht.GenMap()

	// val := []string{}

}

func TestTreeGen(t *testing.T) {}
