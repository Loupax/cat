// The purpose of this project is to create a port of Unix cat in golang
// for education purposes.
// This tool is not supposed to replace the `cat` command of your system
// and comes without any guarantees

/*
** 2018 April 15
**
** The author disclaims copyright to this source code.  In place of
** a legal notice, here is a blessing, copied from the SQlite source code:
**
**    May you do good and not evil.
**    May you find forgiveness for yourself and forgive others.
**    May you share freely, never taking more than you give.
**
*************************************************************************/

package main

import "os"
import (
	"bufio"
	"flag"
	"fmt"
	"io"
)

func OutputLines(out io.Writer, reader *bufio.Reader) {
	for {
		text, err := reader.ReadString('\n')
		fmt.Fprint(out, text)
		if err == io.EOF {
			break
		}
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		reader := bufio.NewReader(os.Stdin)
		OutputLines(os.Stdout, reader)
		os.Exit(0)
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			panic(err)
		}

		OutputLines(os.Stdout, bufio.NewReader(file))
		file.Close()
	}

	os.Exit(0)
}
