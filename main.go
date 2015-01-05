package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

var sortable []string
var unsortable []string

func getLines(r io.Reader) []string {
	var lines []string
	var l string
	var err error

	br := bufio.NewReader(r)
	for err != io.EOF {
		l, err = br.ReadString('\n')
		if err != nil && err != io.EOF {
			die(err)
		}

		l = strings.TrimRight(l, "\n")
		lines = append(lines, l)
	}
	return lines
}

func sortLines(lines []string) []string {
	sl := make([]string, 0, len(lines))
	ul := make([]string, 0, len(lines))

	for _, l := range lines {
		switch {
		case strings.HasPrefix(l, "?   \t"):
			ul = append(ul, l)
		case strings.HasPrefix(l, "ok  \t"):
			sl = append(sl, l)
		default: // garbage. discard it.
		}
	}

	sort.Sort(sort.Reverse(TestLines(sl))) // reversed to but highest first.
	sl = append(sl, ul...)
	return sl
}

func usage() {
	fmt.Printf("%s: sorts `go test` output by duration\n", os.Args[0])
	fmt.Printf("usage: go test ./... | %s\n", os.Args[0])
}

func main() {
	if len(os.Args) > 1 {
		usage()
		os.Exit(0)
	}

	lines := getLines(os.Stdin)
	lines = sortLines(lines)
	for _, l := range lines {
		fmt.Println(l)
	}
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type TestLines []string

func (p TestLines) Len() int      { return len(p) }
func (p TestLines) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p TestLines) Less(i, j int) bool {
	in := p.Duration(i)
	jn := p.Duration(j)
	return in < jn
}

func (p TestLines) Duration(i int) time.Duration {
	d, err := TestLineDuration(p[i])
	if err != nil {
		die(err)
	}
	return d
}

func TestLineDuration(l string) (time.Duration, error) {
	tsi := strings.LastIndex(l, "\t")
	if tsi < 0 {
		return 0, fmt.Errorf("invalid test line: %s", l)
	}
	ts := l[tsi+1:]

	return time.ParseDuration(ts)
}
