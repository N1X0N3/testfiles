package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	testCases := map[string][]string{
		"case1.txt":  {"hello"},
		"case2.txt":  {"HELLO"},
		"case3.txt":  {"HeLlo HuMaN"},
		"case4.txt":  {"1Hello 2There"},
		"case5.txt":  {"Hello\\nThere"},
		"case6.txt":  {"{Hello & There #}"},
		"case7.txt":  {"hello There 1 to 2!"},
		"case8.txt":  {"MaD3IrA&LiSboN"},
		"case9.txt":  {"1a\"#FdwHywR&/()="},
		"case10.txt": {"{|}~"},
		"case11.txt": {"[\\]^_ 'a"},
		"case12.txt": {"RGB"},
		"case13.txt": {":;<=>?@"},
		"case14.txt": {"\\!\" #$%&'()*+,-./"},
		"case15.txt": {"ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		"case16.txt": {"abcdefghijklmnopqrstuvwxyz"},
	}

	for input := range testCases {
		if output, err := ioutil.ReadFile("./testCases/" + input); err != nil {
			panic(err)
		} else {
			testCases[input] = append(testCases[input], string(output))
		}
	}

	for each := range testCases {
		os.Args = []string{"cmd", testCases[each][0]}
		got := AsciiArt(false)

		if string(got) != testCases[each][1] {
			t.Errorf("Test failed, expected:\n%s\ngot:\n%s\n", testCases[each][1], string(got))
		}
	}
}
