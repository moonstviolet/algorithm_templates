package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func Scan(a ...interface{}) { fmt.Fscan(in, a...) }
func Scanln(s *string) {
	for *s, _ = in.ReadString('\n'); *s == "\n"; {
		*s, _ = in.ReadString('\n')
	}
	*s = (*s)[:len(*s)-1]
}
func Print(a ...interface{})                 { fmt.Fprint(out, a...) }
func Printf(format string, a ...interface{}) { fmt.Fprintf(out, format, a...) }
func Log(a ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fmt.Fprint(out, "log ", runtime.FuncForPC(pc).Name(), ":", line, " ")
	fmt.Fprint(out, a...)
	fmt.Fprint(out, "\n")
}

func solve() {

}

func main() {
	defer out.Flush()

	T := 1
	for test := 0; test < T; test++ {
		solve()
	}
}
