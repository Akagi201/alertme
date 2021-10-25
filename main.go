package main

import (
	"fmt"
	"os"
	"os/exec"
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
	if len(os.Args[1:]) == 0 {
		fmt.Println("alertme: Alert me when cmd done!")
		fmt.Println("Usage  : alertme [command] [optional parameters to command]")
		os.Exit(0)
	}

	start := time.Now()
	defer func() {
		end := time.Now()
		diff := end.Sub(start)
		// if diff > 30*time.Second {
		beeep.Notify(shortPwd(), fmt.Sprintf(contentFmt, prettyArray(os.Args[1:]), diff), "")
		// }
	}()

	var cmd = exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("%v failed, err: %v", prettyArray(os.Args[1:]), err)
	}
}
