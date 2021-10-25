package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

var contentFmt = `%v done!
duration: %v`

func prettyArray(s []string) string {
	res := ""
	for _, v := range s {
		res += v
		res += " "
	}
	return res
}

func shortPwd() string {
	pwd, _ := os.Getwd()
	s := strings.Split(pwd, "/")
	if len(s) > 3 {
		return strings.Join(s[len(s)-3:len(s)], "/")
	} else {
		return pwd
	}
}

func main() {
	start := time.Now()
	defer func() {
		end := time.Now()
		diff := end.Sub(start)
		beeep.Notify(shortPwd(), fmt.Sprintf(contentFmt, prettyArray(os.Args[1:]), diff), "")
	}()

	// Redirect stdin to stdout
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	}
}
