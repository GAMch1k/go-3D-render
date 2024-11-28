package utils

import (
	"os"
	"os/exec"
)

func KeyReader() *chan string {
	ch := make(chan string)
    go func(ch chan string) {
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        var b []byte = make([]byte, 1)
        for {
            os.Stdin.Read(b)
            ch <- string(b)
        }
    }(ch)
	return &ch
}