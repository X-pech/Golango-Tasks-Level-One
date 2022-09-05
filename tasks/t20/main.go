package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wbtechl1tasks/taskutils"
)

func reverseWords(s *string) string {
	// simply split the line and then reverse the slice and join it
	r := strings.Split(*s, " ")
	taskutils.ReverseSlice(r)
	return strings.Join(r, " ")
}

func main() {
	// to read the whole line we need to use buffered reader
	// and then ReadString from it to newline
	br := bufio.NewReader(os.Stdin)
	line, _ := br.ReadString('\n')
	// line = strings.TrimSuffix(line, "\n") // to not store \n in memory
	line = strings.TrimRight(line, "\n") // to not copy the whole line
	fmt.Println(reverseWords(&line))
}
