package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGoReloaded(t *testing.T) {
	inputFile := "sample.txt"
	outputFile := "result.txt"
	args := []string{"cmd", inputFile, outputFile}

	testCases := [][]string{
		{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Don’t be sad ,because sad backwards is das . And das not good",
			"Don’t be sad, because sad backwards is das. And das not good"},
		{"harold wilson (cap, 2) : ‘ I’m a optimist ,but a optimist who carries a raincoat . ‘",
			"Harold Wilson: ‘I’m an optimist, but an optimist who carries a raincoat.’"},
	}

	for _, each := range testCases {
		if err := ioutil.WriteFile(inputFile, []byte(each[0]), os.ModePerm); err != nil {
			panic(err)
		}

		os.Args = args
		GoReloaded()

		if actual, err := ioutil.ReadFile(outputFile); err != nil {
			panic(err)
		} else if string(actual) != each[1] {
			t.Errorf("Test failed, expected \"%s\", got \"%s\"", each[1], string(actual))
		}
	}
}
