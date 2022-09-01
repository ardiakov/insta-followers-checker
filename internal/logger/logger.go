package logger

import (
	"fmt"
	"os"
)

func StdOut(mes string) {
	fmt.Fprintf(os.Stdout, mes+"\n")
}

func StdErr(t string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s \n", t, err)
}
