package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `１行め
２行め
３行め`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
