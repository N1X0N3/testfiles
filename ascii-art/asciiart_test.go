package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
	"sync"
	"testing"
)

/*	This test file tests the ascii-art project against the first 16 test cases on
	audit page	*/
func TestAsciiArt(t *testing.T) {
	/*	Each key of the map testCases contains the name of the file containing the
		expected out put for each test case, the value for each key is a slice of
		strings, the first element contains the string to be given as an argument
		at run time, the second will contain the string equivalent of expected
		output	*/
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

	/*	Iterating over each of the keys in testCases, reading and storing the
		contents of the corresponding files for each key	*/
	for input := range testCases {
		if output, err := ioutil.ReadFile("./test-cases/" + input); err != nil {
			panic(err)
		} else {
			testCases[input] = append(testCases[input], string(output))
		}
	}

	/*	Iterate through each test case and starting a goroutine for each, this
		is done so instead of waiting for the previous test to complete they can
		all be checked simulaneously	*/
	var wg sync.WaitGroup
	for _, testCase := range testCases {
		wg.Add(1)
		go func(current []string, w *sync.WaitGroup, ti *testing.T) {
			defer w.Done()
			result := getResult(current)
			/*	Fails the project if the test cases expected output doesn't match
				the actual output	*/
			if string(result) != current[1] {
				ti.Errorf("\nTest fails when given the test case:\n\t\"%s\","+
					"\nexpected:\n%s\ngot:\n%s\n\n",
					current[0], current[1], string(result))
			}
		}(testCase, &wg, t)
	}
	wg.Wait()
}

/*	This function imitates the running of "go run . string", which it then pipes
	into a second function "cat -e" to immitate and then returns the result	*/
func getResult(testCase []string) string {
	first := exec.Command("go", "run", ".", testCase[0])
	second := exec.Command("cat", "-e")
	reader, writer := io.Pipe()
	first.Stdout = writer
	second.Stdin = reader
	var buffer bytes.Buffer
	second.Stdout = &buffer
	first.Start()
	second.Start()
	first.Wait()
	writer.Close()
	second.Wait()
	return buffer.String()
}
