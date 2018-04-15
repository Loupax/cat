package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestOutputLines(t *testing.T) {
	// It must output it's input as-is
	reader := bufio.NewReader(strings.NewReader("test\nline two"))
	output := bytes.NewBufferString("")
	OutputLines(output, reader)

	if output.String() != "test\nline two" {
		t.Fatal("OutputLines must output the data inside it's passed reader as is")
	}
}
